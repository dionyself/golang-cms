package main

import (
	"github.com/astaxie/beego"
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
	utils.SessionInit(curr_env)
}

func main() {
	beego.Run()
}
