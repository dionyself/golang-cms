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
	Pool       map[string]orm.Ormer
	Replicated bool
}

func (db *DB) GetOrm(db_name string) orm.Ormer {
	if (db_name == "") || (db_name == "master") {
		db_name = "default"
	}
	_, ok := db.Pool[db_name]
	if ok {
		return db.Pool[db_name]
	}
	db.Pool[db_name] = orm.NewOrmUsingDB(db_name)
	return db.Pool[db_name]
}

func init() {
	fmt.Println("loading utils.db")
	env, _ := web.AppConfig.String("RunMode")
	dbBlk := "databaseConfig-" + env + "::"
	MasterAddress := ""
	SlaveAddress := ""
	masterServerPort := ""
	slaveServerPort := ""
	replicated := false
	Engine, _ := web.AppConfig.String("DatabaseProvider")
	replicated, _ = web.AppConfig.Bool(dbBlk + "replicated")
	masterServerPort, _ = web.AppConfig.String(dbBlk + "masterServerPort")
	slaveServerPort, _ = web.AppConfig.String(dbBlk + "slaveServerPort")
	Username, _ := web.AppConfig.String(dbBlk + "databaseUser")
	UserPassword, _ := web.AppConfig.String(dbBlk + "userPassword")
	MasterServer, _ := web.AppConfig.String(dbBlk + "masterServer")
	SlaveServer, _ := web.AppConfig.String(dbBlk + "slaveServer")
	Name, _ := web.AppConfig.String(dbBlk + "databaseName")
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
		fl_c, _ := web.AppConfig.String(dbBlk + "sqliteFile")
		MasterAddress = "file:" + fl_c
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
		orm.MaxIdleConnections(maxIdle),
		orm.MaxOpenConnections(maxConn))
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
			orm.MaxIdleConnections(maxIdle),
			orm.MaxOpenConnections(maxConn))
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
	MainDatabase = DB{}
	MainDatabase.Pool = make(map[string]orm.Ormer)
	MainDatabase.Replicated = (replicated == true && Engine != "sqlite3")

	insertDemo, _ := web.AppConfig.Bool(dbBlk + "insertDemoData")
	if force && insertDemo {
		InsertDemoData()
	}

	if MainDatabase.Replicated == true {
		MainDatabase.GetOrm("slave")
		MainDatabase.Pool["slave"].Raw("start slave")
	}
}
