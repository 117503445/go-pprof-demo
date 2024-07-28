package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // 引入 pprof
)

func main() {
	// 设置 HTTP 路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// 启动 HTTP 服务器
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("HTTP server failed: %s", err)
	}
}
