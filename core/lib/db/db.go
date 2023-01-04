package db

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// MainDatabase main db
var MainDatabase DB

// DB ...
type DB struct {
	Orm        orm.Ormer
	Replicated bool
}

func init() {
	fmt.Println("loading utils.db")
	env := web.AppConfig.String("RunMode")
	dbBlk := "databaseConfig-" + env + "::"
	MasterAddress := ""
	SlaveAddress := ""
	masterServerPort := ""
	slaveServerPort := ""
	replicated := false
	Engine := web.AppConfig.String("DatabaseProvider")
	replicated, _ = web.AppConfig.Bool(dbBlk + "replicated")
	masterServerPort = web.AppConfig.String(dbBlk + "masterServerPort")
	slaveServerPort = web.AppConfig.String(dbBlk + "slaveServerPort")
	Username := web.AppConfig.String(dbBlk + "databaseUser")
	UserPassword := web.AppConfig.String(dbBlk + "userPassword")
	MasterServer := web.AppConfig.String(dbBlk + "masterServer")
	SlaveServer := web.AppConfig.String(dbBlk + "slaveServer")
	Name := web.AppConfig.String(dbBlk + "databaseName")
	maxIdle := 300
	maxConn := 300
	if Engine == "" {
		panic("Engine not configured valid options: sqlite3, mysql or postgres")
	}
	if Engine != "sqlite3" && MasterServer == "" {
		panic("masterServer not configured")
	}
	if replicated == true && SlaveServer == "" {
		panic("DB Replication: slaveServer not configured")
	}
	if Engine == "mysql" {
		orm.RegisterDriver(Engine, orm.DRMySQL)
		if masterServerPort == "0" {
			masterServerPort = "3306"
		}
		if slaveServerPort == "0" {
			slaveServerPort = "3306"
		}
		MasterAddress = Username + ":" + UserPassword + "@tcp(" + MasterServer + ":" + masterServerPort + ")/" + Name + "?charset=utf8"
		if replicated == true {
			SlaveAddress = Username + ":" + UserPassword + "@tcp(" + SlaveServer + ":" + slaveServerPort + ")/" + Name + "?charset=utf8"
		}
	} else if Engine == "sqlite3" {
		orm.RegisterDriver(Engine, orm.DRSqlite)
		MasterAddress = "file:" + web.AppConfig.String(dbBlk+"sqliteFile")
	} else if Engine == "postgres" {
		orm.RegisterDriver(Engine, orm.DRPostgres)
		if masterServerPort == "0" {
			masterServerPort = "5432"
		}
		if slaveServerPort == "0" {
			slaveServerPort = "5432"
		}
		MasterAddress = "user=" + Username + " password=" + UserPassword + " dbname=" + Name + " host=" + MasterServer + " port=" + masterServerPort + " sslmode=disable"
		if replicated == true {
			SlaveAddress = "user=" + Username + " password=" + UserPassword + " dbname=" + Name + " host=" + SlaveServer + " port=" + slaveServerPort + " sslmode=disable"
		}
	}
	err := orm.RegisterDataBase(
		"default",
		Engine,
		MasterAddress,
		maxIdle,
		maxConn)
	if err != nil {
		panic("DB: cannot register DB on master")
	} else if replicated == true && Engine != "sqlite3" {
		if MasterAddress == SlaveAddress {
			panic("DB Replication: MasterAddress and SlaveAddress are equal!")
		}
		err = orm.RegisterDataBase(
			"slave",
			Engine,
			SlaveAddress,
			maxIdle,
			maxConn)
	}

	// DB SETUP
	orm.Debug, _ = web.AppConfig.Bool("DatabaseDebugMode")
	force, _ := web.AppConfig.Bool("ReCreateDatabase")
	verbose, _ := web.AppConfig.Bool("DatabaseLogging")
	err = orm.RunSyncdb("default", force, verbose)
	if err != nil {
		panic(err)
	} else if replicated == true && Engine != "sqlite3" {
		err = orm.RunSyncdb("slave", force, verbose)
	}
	if err != nil {
		panic(err)
	}

	MainDatabase.Orm = orm.NewOrm()
	MainDatabase.Replicated = (replicated == true && Engine != "sqlite3")

	insertDemo, _ := web.AppConfig.Bool(dbBlk + "insertDemoData")
	if force && insertDemo {
		InsertDemoData()
	}

	if MainDatabase.Replicated == true {
		MainDatabase.Orm.Using("slave")
		MainDatabase.Orm.Raw("start slave")
	}
}
