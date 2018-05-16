package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	// 注意斜杠！
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/"+"download/")))) // 正确
	//http.Handle("/", http.FileServer(http.Dir("public"))) // 正确（访问根目录时转到public目录）
        fmt.Println("request in")
	//http.Handle("/public", http.StripPrefix("/public", http.FileServer(http.Dir("public")))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("/public/"))) // 错误
	//http.Handle("/public", http.FileServer(http.Dir("./public"))) // 错误！

	log.Fatal(http.ListenAndServe(":8089", nil))
}