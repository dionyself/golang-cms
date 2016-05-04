package utils

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var Cache CACHE

type CACHE struct {
	isEnabled bool
	servers   map[string]redis.Conn
	mode      string
}

func (cache *CACHE) GetString(origin string, commandName string, args ...interface{}) (reply string, found bool) {
	if !cache.isAvailable() {
		return "", false
	}
	value, err := redis.String(cache.servers[origin].Do(commandName, args...))
	if err != nil {
		if cache.mayFallback(origin) {
			value, err = redis.String(cache.servers["master"].Do(commandName, args...))
		} else {
			return "", false
		}
	}
	return value, err == nil
}

func (cache *CACHE) isAvailable() bool {
	if cache.isEnabled == true && cache.servers["slave"] != nil {
		return true
	}
	return false
}

func (cache *CACHE) mayFallback(origin string) bool {
	if origin == "slave" && cache.servers["master"] != cache.servers[origin] {
		return true
	}
	return false
}

func init() {
	Cache = CACHE{isEnabled: false}
	fmt.Println("loaded utils.Cache")
}
