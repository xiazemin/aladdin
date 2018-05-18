package main

import (
"fmt"
"io"
"net/http"
"os"
)

const (
upload_path string = "/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/test/"
)

func load_success(w http.ResponseWriter, r *http.Request) {
io.WriteString(w, "上传成功!")

}

//上传
func uploadHandle(w http.ResponseWriter, r *http.Request) {
//从请求当中判断方法
if r.Method == "GET" {
io.WriteString(w, "<html><head><title>上传</title></head>"+
"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
"<label>上传日志</label>"+":"+
"<input type=\"file\" name='file'  /><br/><br/>    "+
"<label><input type=\"submit\" value=\"上传日志\"/></label></form></body></html>")
} else {
//获取文件内容 要这样获取
file, head, err := r.FormFile("file")
if err != nil {
fmt.Println(err)
return
}
defer file.Close()
//创建文件
fW, err := os.Create(upload_path + head.Filename)
if err != nil {
fmt.Println("文件创建失败")
return
}
defer fW.Close()
_, err = io.Copy(fW, file)
if err != nil {
fmt.Println("文件保存失败")
return
}
//io.WriteString(w, head.Filename+" 保存成功")
http.Redirect(w, r, "/success", http.StatusFound)
//io.WriteString(w, head.Filename)
}
}

func main() {
fmt.Println("OK!")
//启动一个http 服务器
http.HandleFunc("/success", load_success)
//上传
http.HandleFunc("/upload", uploadHandle)
err := http.ListenAndServe("127.0.0.1:8089", nil)
if err != nil {
fmt.Println("服务器启动失败")
return
}
fmt.Println("服务器启动成功")
}
