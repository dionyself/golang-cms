package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/dionyself/golang-cms/models"
	_ "github.com/dionyself/golang-cms/routers"
)

func init() {
	DbEngine := beego.AppConfig.String("DB_Engine")
	if DbEngine == "mysql" {
		orm.RegisterDriver(DbEngine, orm.DR_MySQL)
	}
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase(
		"default",
		DbEngine,
		beego.AppConfig.String("DB_Username")+":"+
			beego.AppConfig.String("DB_UserPassword")+"@/"+
			beego.AppConfig.String("DB_Name")+"?charset=utf8",
		maxIdle,
		maxConn)
}

func main() {
	beego.SessionOn = true
	// DB SETUP
	name := "default"
	orm.Debug, _ = beego.AppConfig.Bool("DebugMode")
	force, _ := beego.AppConfig.Bool("ReCreateDB")
	verbose := true // Print log.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
