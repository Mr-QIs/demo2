package main

import (
	"fmt"
	"log"
	"net/http"
)

// homeHandler 处理根路径请求
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// main 主函数
func main() {
	// 设置静态文件服务
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 设置路由
	http.HandleFunc("/", homeHandler)

	// 启动服务器
	port := ":8080"
	fmt.Printf("服务器正在运行，访问 http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
