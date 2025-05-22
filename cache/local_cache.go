package cache

import (
	"log"
	"sync"

	"github.com/bluele/gcache"
)

// localCache 本地内存缓存实现
type localCache struct {
	mutex sync.RWMutex
	cache gcache.Cache
}

func newLocalCache(size int) *localCache {
	return &localCache{
		cache: gcache.New(size).LRU().Build(),
	}
}

func (c *localCache) Get(key string) (any, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, err := c.cache.Get(key)
	if err != nil {
		log.Println("获取缓存数据失败：", err)
		return nil, false
	}
	return value, true
}

func (c *localCache) Set(key string, val any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Set(key, val)
}

func (c *localCache) Close() error {
	c.cache.Purge()
	return nil
}
