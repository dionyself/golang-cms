package main

import (
	"github.com/dionyself/beego"
	_ "github.com/dionyself/beego/session/redis"
	_ "github.com/dionyself/golang-cms/core/template"
	_ "github.com/dionyself/golang-cms/routers"
	"github.com/dionyself/golang-cms/utils"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	currentEnvironment := beego.AppConfig.String("RunMode")
	utils.SessionInit(currentEnvironment)
}

func main() {
	beego.Run()
}
