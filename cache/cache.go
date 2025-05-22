package cache

import (
	"img2color/util"
	"log"
	"strconv"
)

type Cache interface {
	Get(key string) (any, bool) // 获取缓存数据
	Set(key string, value any)  // 缓存数据
	Close() error               // 关闭缓存
}

func NewCache() Cache {
	// 缓存内存，默认本地缓存
	cacheType := util.GetEnvDefault("CACHE_TYPE", util.CacheType)
	var cache Cache
	if cacheType == "local" {
		log.Println("使用本地缓存...")
		capacity, _ := strconv.Atoi(util.GetEnvDefault("LOCAL_CACHE_CAPACITY", util.LocalCacheCapacity))
		cache = newLocalCache(capacity)
	} else if cacheType == "redis" {
		log.Println("使用Redis缓存...")
		db, _ := strconv.Atoi(util.GetEnvDefault("REDIS_DB", util.RedisCacheDB)) // ignore error if it doesn't exist
		cache = newRedisCache(
			util.GetEnvDefault("REDIS_ADDR", util.RedisCacheAddr),
			util.GetEnvDefault("REDIS_PASSWORD", util.RedisCachePassword),
			db,
		)
	}

	return cache
}
