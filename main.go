package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/dionyself/golang-cms/routers"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	DbAddress := ""
	DbEngine := beego.AppConfig.String("DB_Engine")
	DbServerPort := beego.AppConfig.String("DB_ServerPort")
	DbUsername := beego.AppConfig.String("DB_Username")
	DbUserPassword := beego.AppConfig.String("DB_UserPassword")
	DbServer := beego.AppConfig.String("DB_Server")
	DbName := beego.AppConfig.String("DB_Name")
	maxIdle := 300
	maxConn := 300
	if DbEngine == "mysql" {
		orm.RegisterDriver(DbEngine, orm.DRMySQL)
		if DbServerPort == "0" {
			DbServerPort = "3306"
		}
		DbAddress = DbUsername + ":" + DbUserPassword + "@tcp(" + DbServer + ":" + DbServerPort + ")/" + DbName + "?charset=utf8"
	} else if DbEngine == "sqlite3" {
		orm.RegisterDriver(DbEngine, orm.DRSqlite)
		DbAddress = "file:" + beego.AppConfig.String("DB_SqliteFile")
	} else if DbEngine == "postgres" {
		orm.RegisterDriver(DbEngine, orm.DRPostgres)
		if DbServerPort == "0" {
			DbServerPort = "5432"
		}
		DbAddress = "user=" + DbUsername + " password=" + DbUserPassword + " dbname=" + DbName + " host=" + DbServer + " port=" + DbServerPort + " sslmode=disable"
	}
	err := orm.RegisterDataBase(
		"default",
		DbEngine,
		DbAddress,
		maxIdle,
		maxConn)
	if err != nil {
		panic(err)
	}
}

func main() {
	// DB SETUP
	orm.Debug, _ = beego.AppConfig.Bool("DB_DebugMode")
	force, _ := beego.AppConfig.Bool("DB_ReCreate")
	verbose, _ := beego.AppConfig.Bool("DB_Logging")
	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
