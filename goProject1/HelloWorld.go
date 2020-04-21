package main

import (
	"fmt"
	"net/http"
)

func main() {
	//打印
	fmt.Println("Hello World")
	/*搭建一个http服务，实现基础打印*/
	//指定当前目录为根目录
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", helloWorld)
	//使用端口8080作为服务监控端口
	http.ListenAndServe(":8080", nil)
}

//在浏览器中打印Hello World
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
