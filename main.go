package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/dionyself/golang-cms/routers"
	"github.com/dionyself/golang-cms/utils"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = `127.0.0.2:6379,100`
	curr_env := beego.AppConfig.String("RunMode")
	err := utils.DatabaseInit(curr_env)
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
	insertDemo, _ := beego.AppConfig.Bool("DB_InsertDemoData")
	if force && insertDemo {
		utils.InsertDemoData()
	}
	beego.Run()
}
