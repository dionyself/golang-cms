package utils

import (
	_ "fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func SessionInit(env string) error {
  sessBlk := "sessionConfig-" + env + "::"
	provider := beego.AppConfig.String("SessionProvider")
	Address := ""
	if provider == "redis"{
		Address = beego.AppConfig.String(sessBlk+"redisServer")
		if Address == "" {
			Address = beego.AppConfig.String("cacheConfig-"+env+"::redisMasterServer")
		}
	}

  //sessionConfig = "{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": Address}"

		beego.BConfig.WebConfig.Session.SessionName = beego.AppConfig.String(sessBlk+"cookieName")
    beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = beego.AppConfig.String(sessBlk+"gclifetime")
					"providerConfig":  filepath.ToSlash(BConfig.WebConfig.Session.SessionProviderConfig),
					"secure":          BConfig.Listen.EnableHTTPS,
					"enableSetCookie": BConfig.WebConfig.Session.SessionAutoSetCookie,
					"domain":          BConfig.WebConfig.Session.SessionDomain,
					"cookieLifeTime":  BConfig.WebConfig.Session.SessionCookieLifeTime,
}
