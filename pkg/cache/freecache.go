package cache

import (
	"github.com/coocood/freecache"
)

var cache = freecache.NewCache(100 * 1024 * 1024)

func Set(key, value []byte, expireSeconds int) error {
	return cache.Set(key, value, expireSeconds)
}

func Get(key []byte) ([]byte, error) {
	return cache.Get(key)
}

func Del(key []byte) bool {
	return cache.Del(key)
}
