//HTTP协议通常承载于TCP协议之上
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//注册默认路由
	// /go 是请求路径
	//myHandler是自己写的回调函数
	http.HandleFunc("/go", myHandler)
	//addr：监听的地址
	//hander：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// hander函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	//请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	//回复
	w.Write([]byte("www.51mh.com"))
}
