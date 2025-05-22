package util

import "os"

const (
	CacheType = "local"
	// 本地缓存的默认配置
	LocalCacheCapacity = "100000"
	// Redis缓存的默认配置
	RedisCacheAddr     = "127.0.0.1:6379"
	RedisCachePassword = ""
	RedisCacheDB       = "0"
	AllowReffers       = "*"
)

// 获取环境变量，如果不存在则使用默认值
func GetEnvDefault(key, defaultValue string) string {
	value, isExists := os.LookupEnv(key)
	if !isExists {
		value = defaultValue
	}
	return value
}
