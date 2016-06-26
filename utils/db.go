package utils

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Mdb DB

type DB struct {
	Orm        orm.Ormer
	replicated bool
}

func init() {
	fmt.Println("loading utils.db")
	env := beego.AppConfig.String("RunMode")
	dbBlk := "databaseConfig-" + env + "::"
	MasterAddress := ""
	SlaveAddress := ""
	replicated := false
	Engine := beego.AppConfig.String("DatabaseProvider")
	replicated, _ = beego.AppConfig.Bool(dbBlk + "replicated")
	ServerPort := beego.AppConfig.String(dbBlk + "serverPort")
	Username := beego.AppConfig.String(dbBlk + "databaseUser")
	UserPassword := beego.AppConfig.String(dbBlk + "userPassword")
	MasterServer := beego.AppConfig.String(dbBlk + "masterServer")
	SlaveServer := beego.AppConfig.String(dbBlk + "slaveServer")
	Name := beego.AppConfig.String(dbBlk + "databaseName")
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
		if ServerPort == "0" {
			ServerPort = "3306"
		}
		MasterAddress = Username + ":" + UserPassword + "@tcp(" + MasterServer + ":" + ServerPort + ")/" + Name + "?charset=utf8"
		if replicated == true {
			SlaveAddress = Username + ":" + UserPassword + "@tcp(" + SlaveServer + ":" + ServerPort + ")/" + Name + "?charset=utf8"
		}
	} else if Engine == "sqlite3" {
		orm.RegisterDriver(Engine, orm.DRSqlite)
		MasterAddress = "file:" + beego.AppConfig.String(dbBlk+"sqliteFile")
	} else if Engine == "postgres" {
		orm.RegisterDriver(Engine, orm.DRPostgres)
		if ServerPort == "0" {
			ServerPort = "5432"
		}
		MasterAddress = "user=" + Username + " password=" + UserPassword + " dbname=" + Name + " host=" + MasterServer + " port=" + ServerPort + " sslmode=disable"
		if replicated == true {
			SlaveAddress = "user=" + Username + " password=" + UserPassword + " dbname=" + Name + " host=" + SlaveServer + " port=" + ServerPort + " sslmode=disable"
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
	orm.Debug, _ = beego.AppConfig.Bool("DatabaseDebugMode")
	force, _ := beego.AppConfig.Bool("ReCreateDatabase")
	verbose, _ := beego.AppConfig.Bool("DatabaseLogging")
	err = orm.RunSyncdb("default", force, verbose)
	if err != nil {
		panic(err)
	} else if replicated == true && Engine != "sqlite3" {
		err = orm.RunSyncdb("slave", force, verbose)
	}
	if err != nil {
		panic(err)
	}

	Mdb.Orm = orm.NewOrm()
	Mdb.replicated = (replicated == true && Engine != "sqlite3")

	insertDemo, _ := beego.AppConfig.Bool(dbBlk + "insertDemoData")
	if force && insertDemo {
		InsertDemoData()
	}

	if Mdb.replicated == true {
		Mdb.Orm.Using("slave")
		Mdb.Orm.Raw("start slave")
	}
}
