package utils

import (
	"github.com/beego/beego/v2/server/web"
)

// SessionInit ...
func SessionInit(env string) {
	sessBlk := "sessionConfig-" + env + "::"
	provider := web.AppConfig.String("SessionProvider")
	Address := ""
	if provider == "redis" {
		Address = web.AppConfig.String(sessBlk + "redisServer")
		if Address == "" {
			Address = web.AppConfig.String("cacheConfig-" + env + "::redisMasterServer")
		}
	}

	web.BConfig.WebConfig.Session.SessionName = web.AppConfig.String(sessBlk + "cookieName")
	web.BConfig.WebConfig.Session.SessionGCMaxLifetime, _ = web.AppConfig.Int64(sessBlk + "gclifetime")
	web.BConfig.WebConfig.Session.SessionProviderConfig = Address
	web.BConfig.Listen.EnableHTTPS, _ = web.AppConfig.Bool(sessBlk + "secure")
	web.BConfig.WebConfig.Session.SessionAutoSetCookie, _ = web.AppConfig.Bool(sessBlk + "enableSetCookie")
	web.BConfig.WebConfig.Session.SessionDomain = web.AppConfig.String(sessBlk + "domain")
	web.BConfig.WebConfig.Session.SessionCookieLifeTime, _ = web.AppConfig.Int(sessBlk + "cookieLifeTime")
}
