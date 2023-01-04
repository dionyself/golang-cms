package main

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	_ "github.com/dionyself/golang-cms/core/template"
	_ "github.com/dionyself/golang-cms/routers"
	"github.com/dionyself/golang-cms/utils"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	currentEnvironment := config.String("RunMode")
	utils.SessionInit(currentEnvironment)
}

func main() {
	web.Run()
}
