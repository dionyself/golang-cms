package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/dionyself/golang-cms/models"
	_ "github.com/dionyself/golang-cms/routers"
)

//var globalSessions *session.Manager

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "golang_cms:golang_cms@/golang_cms?charset=utf8", maxIdle, maxConn)
}

func main() {
	beego.SessionOn = true

	// DB SETUP
	name := "default"
	orm.Debug = true
	force := true   // re-create db
	verbose := true // Print log.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
