package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"img2color/cache"
	"log"
	"net/http"
	"sync"
)

var (
	imageCache cache.Cache
	once       sync.Once
)

// 获取缓存实例，确保只初始化一次
func getCache() cache.Cache {
	once.Do(func() {
		imageCache = cache.NewCache()
	})
	return imageCache
}

// 计算图片URL的哈希值作为缓存键
func hashImgURL(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])
}

// CacheMiddleware 缓存中间件，对图片颜色提取结果进行缓存
func CacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取图片URL参数
		imgURL := r.URL.Query().Get("img")
		if imgURL == "" {
			next(w, r)
			return
		}

		// 计算缓存键
		cacheKey := hashImgURL(imgURL)

		// 尝试从缓存获取结果
		cache := getCache()
		if cachedData, found := cache.Get(cacheKey); found {
			if data, ok := cachedData.(string); ok {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-Cache", "HIT")
				w.Write([]byte(data))
				log.Println("Cache hit for:", imgURL)
				return
			}
		}

		// 创建一个响应记录器来捕获响应
		rw := newResponseWriter(w)

		// 调用下一个处理函数
		next(rw, r)

		// 如果状态码是200，则缓存响应
		if rw.statusCode == http.StatusOK {
			cache.Set(cacheKey, string(rw.body))
			w.Header().Set("X-Cache", "MISS")
		}
	}
}

// responseWriter 是一个自定义的响应写入器，用于捕获响应内容
type responseWriter struct {
	http.ResponseWriter
	body       []byte
	statusCode int
}

// 创建一个新的响应记录器
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

// Write 实现http.ResponseWriter接口
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}

// WriteHeader 实现http.ResponseWriter接口
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
