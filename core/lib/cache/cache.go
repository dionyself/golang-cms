package cache

import (
	"time"

	"github.com/astaxie/beego"
	redisCache "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	internalCache "github.com/patrickmn/go-cache"
)

// MainCache is global cache
var MainCache CACHE

// CACHE main ...
type CACHE struct {
	isEnabled         bool
	servers           map[string]redisCache.Cache
	internal          *internalCache.Cache
	DefaultExpiration time.Duration
	dualmode          bool
}

// GetString bring string from cache, expiration time in seconds
func (cache *CACHE) GetString(cacheKey string, expirationTime int64) (string, bool) {
	if !cache.isEnabled {
		return "", false
	}
	payload, found := cache.internal.Get(cacheKey)
	if found {
		result, err := redis.String(payload, nil)
		if err == nil {
			return result, true
		}
	}
	if cache.dualmode {
		server := cache.servers["slave"]
		result, err := redis.String(server.Get(cacheKey), nil)
		if err == nil {
			cache.Set(cacheKey, result, expirationTime)
			return result, true
		}
	}
	return "", false
}

// Set any value to cache, expiration time in seconds
func (cache *CACHE) Set(cacheKey string, value interface{}, expirationTime int64) bool {
	if !cache.isEnabled {
		return false
	}
	duration := cache.DefaultExpiration
	if expirationTime != 0 {
		duration = time.Duration(expirationTime) * time.Second
	} else {
		duration = cache.DefaultExpiration
	}
	cache.internal.Set(cacheKey, value, duration)
	if cache.dualmode {
		server := cache.servers["master"]
		err := server.Put(cacheKey, value, duration)
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
	MainCache = CACHE{isEnabled: isEnable, dualmode: dualmode}
	MainCache.internal = internalCache.New(time.Duration(defaultExpiry)*time.Second, time.Duration(flushInterval)*time.Second)
	MainCache.DefaultExpiration = internalCache.DefaultExpiration
	MainCache.servers = make(map[string]redisCache.Cache)
	if dualmode {
		masterRedis := redisCache.Cache{}
		slaveRedis := redisCache.Cache{}
		_ = masterRedis.StartAndGC(master)
		_ = slaveRedis.StartAndGC(slave)
		MainCache.servers["slave"] = slaveRedis
		MainCache.servers["master"] = masterRedis
	}
}
