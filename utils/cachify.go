package utils

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	Rcache "github.com/astaxie/beego/cache"
	"github.com/garyburd/redigo/redis"
	Lcache "github.com/patrickmn/go-cache"
)

var Mcache CACHE

type CACHE struct {
	isEnabled         bool
	servers           map[string]Rcache.Cache
	internal          *Lcache.Cache
	DefaultExpiration time.Duration
	dualmode          bool
}

func (cache *CACHE) GetString(key string, expire int64) (string, bool) {
	if !cache.isEnabled {
		return "", false
	}
	payload, found := cache.internal.Get(key)
	if found {
		result, err := redis.String(payload, nil)
		if err == nil {
			return result, true
		}
	}
	if cache.dualmode {
		server := cache.servers["slave"]
		result, err := redis.String(server.Get(key), nil)
		if err == nil {
			cache.Set(key, result, expire)
			return result, true
		}
	}
	return "", false
}

func (cache *CACHE) Set(key string, value interface{}, expire int64) bool {
	if !cache.isEnabled {
		return false
	}
	duration := cache.DefaultExpiration
	if expire != 0 {
		duration = time.Duration(expire) * time.Second
	} else {
		duration = cache.DefaultExpiration
	}
	cache.internal.Set(key, value, duration)
	if cache.dualmode {
		server := cache.servers["master"]
		err := server.Put(key, value, duration)
		if err == nil {
			return true
		}
	}
	return true
}

func init() {
	env := beego.AppConfig.String("RunMode")
	cacheBlk := "cacheConfig-" + env + "::"
	isEnable, _ := beego.AppConfig.Bool(cacheBlk + "enabled")
	dualmode, _ := beego.AppConfig.Bool(cacheBlk + "dualmode")
	master := beego.AppConfig.String(cacheBlk + "redisMasterServer")
	slave := beego.AppConfig.String(cacheBlk + "redisSlaveServer")
	flushInterval, _ := beego.AppConfig.Int64(cacheBlk + "flushInterval")
	defaultExpiry, _ := beego.AppConfig.Int64(cacheBlk + "defaultExpiry")
	Mcache = CACHE{isEnabled: isEnable, dualmode: dualmode}
	Mcache.internal = Lcache.New(time.Duration(defaultExpiry)*time.Second, time.Duration(flushInterval)*time.Second)
	Mcache.DefaultExpiration = Lcache.DefaultExpiration
	Mcache.servers = make(map[string]Rcache.Cache)
	if dualmode {
		Mcache.servers["slave"], _ = Rcache.NewCache("redis", slave)
		Mcache.servers["master"], _ = Rcache.NewCache("redis", master)
	}
	fmt.Println("loaded utils.Cache")
}
