package util

import "os"

const (
	AllowReffers = "*"
)

// 获取环境变量，如果不存在则使用默认值
func GetEnvDefault(key, defaultValue string) string {
	value, isExists := os.LookupEnv(key)
	if !isExists {
		value = defaultValue
	}
	return value
}
