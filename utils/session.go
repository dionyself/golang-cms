package utils

import (
	"github.com/astaxie/beego"
)

// SessionInit ...
func SessionInit(env string) {
	sessBlk := "sessionConfig-" + env + "::"
	provider := beego.AppConfig.String("SessionProvider")
	Address := ""
	if provider == "redis" {
		Address = beego.AppConfig.String(sessBlk + "redisServer")
		if Address == "" {
			Address = beego.AppConfig.String("cacheConfig-" + env + "::redisMasterServer")
		}
	}

	beego.BConfig.WebConfig.Session.SessionName = beego.AppConfig.String(sessBlk + "cookieName")
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime, _ = beego.AppConfig.Int64(sessBlk + "gclifetime")
	beego.BConfig.WebConfig.Session.SessionProviderConfig = Address
	beego.BConfig.Listen.EnableHTTPS, _ = beego.AppConfig.Bool(sessBlk + "secure")
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie, _ = beego.AppConfig.Bool(sessBlk + "enableSetCookie")
	beego.BConfig.WebConfig.Session.SessionDomain = beego.AppConfig.String(sessBlk + "domain")
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime, _ = beego.AppConfig.Int(sessBlk + "cookieLifeTime")
}
