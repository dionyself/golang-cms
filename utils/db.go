package utils

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func DatabaseInit(env string) error {
	dbBlk := "databaseConfig-" + env + "::"
	MasterAddress := ""
	SlaveAddress := ""
	Engine := beego.AppConfig.String("DatabaseProvider")
	ServerPort := beego.AppConfig.String(dbBlk + "serverPort")
	Username := beego.AppConfig.String(dbBlk + "databaseUser")
	UserPassword := beego.AppConfig.String(dbBlk + "userPassword")
	MasterServer := beego.AppConfig.String(dbBlk + "masterServer")
	SlaveServer := beego.AppConfig.String(dbBlk + "slaveServer")
	Name := beego.AppConfig.String(dbBlk + "databaseName")
	maxIdle := 300
	maxConn := 300
	if Engine == "mysql" {
		orm.RegisterDriver(Engine, orm.DRMySQL)
		if ServerPort == "0" {
			ServerPort = "3306"
		}
		MasterAddress = Username + ":" + UserPassword + "@tcp(" + MasterServer + ":" + ServerPort + ")/" + Name + "?charset=utf8"
		if SlaveServer != "" {
			SlaveAddress = Username + ":" + UserPassword + "@tcp(" + SlaveServer + ":" + ServerPort + ")/" + Name + "?charset=utf8"
		} else {
			SlaveAddress = MasterAddress
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
		if SlaveServer != "" {
			SlaveAddress = "user=" + Username + " password=" + UserPassword + " dbname=" + Name + " host=" + SlaveServer + " port=" + ServerPort + " sslmode=disable"
		} else {
			SlaveAddress = MasterAddress
		}
	}
	err := orm.RegisterDataBase(
		"default",
		Engine,
		MasterAddress,
		maxIdle,
		maxConn)
	if err != nil {
		return err
	} else if SlaveAddress != "" {
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
		fmt.Println(err)
	}
	insertDemo, _ := beego.AppConfig.Bool(dbBlk + "insertDemoData")
	if force && insertDemo {
		InsertDemoData()
	}

	return err
}
