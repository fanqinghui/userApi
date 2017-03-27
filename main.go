package main

import (
	"fmt"
	"net/http"
	_ "userApi/routers"
)

//原生go实现方式，没用beego等框架
func main() {
	err := http.ListenAndServe(":8080", nil) //启动监听服务
	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}
