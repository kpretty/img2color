package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// redisCache Redis缓存实现
type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

func newRedisCache(addr, password string, db int) *redisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisCache{
		client: client,
		ctx:    context.Background(),
	}
}

// Get 获取缓存数据
func (c *redisCache) Get(key string) (any, bool) {
	val, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		log.Println("获取缓存数据失败：", err)
		return nil, false
	}

	return val, true
}

func (c *redisCache) Set(key string, value any) {
	c.client.Set(c.ctx, key, value, 0)
}

func (c *redisCache) Close() error {
	return c.client.Close()
}
