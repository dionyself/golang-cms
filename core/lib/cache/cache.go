package cache

import (
	"encoding/json"
	"fmt"
	"time"

	redisCacheAdapter "github.com/beego/beego/v2/adapter/cache"
	redisCache "github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/server/web"
	internalCache "github.com/patrickmn/go-cache"
)

// MainCache is global cache
var MainCache CACHE

// CACHE main ...
type CACHE struct {
	isEnabled         bool
	servers           map[string]redisCacheAdapter.Cache
	internal          *internalCache.Cache
	DefaultExpiration time.Duration
	dualmode          bool
}

// GetMap bring string from cache, expiration time in seconds
func (cache *CACHE) GetMap(cacheKey string, expirationTime int64) (map[string]string, bool) {
	if !cache.isEnabled {
		return nil, false
	}
	var result map[string]string
	Map := func(incomingValue interface{}) map[string]string {
		switch value := incomingValue.(type) {
		case map[string]string:
			result = value
		case []byte:
			var newValue map[string]string
			json.Unmarshal(value, &newValue)
			result = newValue
		case string:
			var newValue map[string]string
			json.Unmarshal([]byte(value), &newValue)
			result = newValue
		default:
			result = nil
		}
		return result
	}
	payload, found := cache.internal.Get(cacheKey)
	if found {
		result = Map(payload)
		if result != nil {
			return result, true
		}
	}
	if cache.dualmode {
		server := cache.servers["slave"]
		slaveResult := server.Get(cacheKey)
		if slaveResult != nil {
			go cache.Set(cacheKey, slaveResult, expirationTime)
			result = Map(slaveResult)
			if result != nil {
				return result, true
			}
		}
	}
	return nil, false
}

func (cache *CACHE) GetStringList(cacheKey string, expirationTime int64) []string {
	return []string{}
}

// GetString bring string from cache, expiration time in seconds
func (cache *CACHE) GetString(cacheKey string, expirationTime int64) (string, bool) {
	if !cache.isEnabled {
		return "", false
	}
	var result string
	String := func(incomingValue interface{}) string {
		switch value := incomingValue.(type) {
		case string:
			result = value
		case int32, int64:
			result = fmt.Sprintf("%v", value)
		case []byte:
			result = string(value[:])
		case map[string]string:
			jsonValue, _ := json.Marshal(value)
			result = string(jsonValue[:])
		default:
			result = ""
		}
		return result
	}
	payload, found := cache.internal.Get(cacheKey)
	if found {
		result = String(payload)
		if result != "" {
			return result, true
		}
	}
	if cache.dualmode {
		server := cache.servers["slave"]
		slaveResult := server.Get(cacheKey)
		if slaveResult != nil {
			go cache.Set(cacheKey, slaveResult, expirationTime)
			result = String(slaveResult)
			if result != "" {
				return result, true
			}
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
	env, _ := web.AppConfig.String("RunMode")
	cacheBlk := "cacheConfig-" + env + "::"
	isEnable, _ := web.AppConfig.Bool(cacheBlk + "enabled")
	dualmode, _ := web.AppConfig.Bool(cacheBlk + "dualmode")
	master, _ := web.AppConfig.String(cacheBlk + "redisMasterServer")
	slave, _ := web.AppConfig.String(cacheBlk + "redisSlaveServer")
	flushInterval, _ := web.AppConfig.Int64(cacheBlk + "flushInterval")
	defaultExpiry, _ := web.AppConfig.Int64(cacheBlk + "defaultExpiry")
	MainCache = CACHE{isEnabled: isEnable, dualmode: dualmode}
	MainCache.internal = internalCache.New(time.Duration(defaultExpiry)*time.Second, time.Duration(flushInterval)*time.Second)
	MainCache.DefaultExpiration = internalCache.DefaultExpiration
	MainCache.servers = make(map[string]redisCacheAdapter.Cache)
	if dualmode {
		masterRedis, _ := redisCache.NewCache("redis", master)
		slaveRedis, _ := redisCache.NewCache("redis", slave)
		MainCache.servers["slave"] = redisCacheAdapter.CreateNewToOldCacheAdapter(slaveRedis)
		MainCache.servers["master"] = redisCacheAdapter.CreateNewToOldCacheAdapter(masterRedis)
	}
}
