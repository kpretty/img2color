package api

import (
	"fmt"

	"encoding/json"
	"img2color/img"
	"img2color/middleware"
	"img2color/util"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var allowedReferers []string

// -- only to vercel start --
func init() {
	allowedReferers = parseReferers(util.GetEnvDefault("ALLOWED_REFERERS", util.AllowReffers))
}

func main() {
	// 使用缓存中间件包装handler函数
	http.HandleFunc("/api", middleware.CacheMiddleware(handler))
	log.Println("服务器监听在: 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动服务器时出错：%v\n", err)
	}
}

// -- only to vercel end --

func StartServer() {
	allowedReferers = parseReferers(util.GetEnvDefault("ALLOWED_REFERERS", util.AllowReffers))
	// 使用缓存中间件包装handler函数
	http.HandleFunc("/api", middleware.CacheMiddleware(handler))
	log.Println("服务器监听在: 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动服务器时出错：%v\n", err)
	}
}

func parseReferers(referers string) []string {
	refererList := strings.Split(referers, ",")
	for i, referer := range refererList {
		refererList[i] = strings.TrimSpace(referer)
	}
	return refererList
}

func handler(w http.ResponseWriter, r *http.Request) {
	handleImageColor(w, r)
}

func isRefererAllowed(referer string) bool {
	if len(allowedReferers) == 0 {
		return true
	}

	for _, allowedReferer := range allowedReferers {
		allowedReferer = strings.ReplaceAll(allowedReferer, ".", "\\.")
		allowedReferer = strings.ReplaceAll(allowedReferer, "*", ".*")
		match, _ := regexp.MatchString(allowedReferer, referer)
		if match {
			return true
		}
	}

	return false
}

func handleImageColor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Referer")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	referer := r.Header.Get("Referer")
	if !isRefererAllowed(referer) {
		http.Error(w, "禁止访问", http.StatusForbidden)
		return
	}

	imgURL := r.URL.Query().Get("img")
	if imgURL == "" {
		http.Error(w, "缺少img参数", http.StatusBadRequest)
		return
	}

	color, err := img.Img2color(imgURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("提取主色调失败：%v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("%s提取主色调成功: %s\n", imgURL, color)

	data := map[string]string{
		"RGB": color,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
