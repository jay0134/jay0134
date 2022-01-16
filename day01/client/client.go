package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回 200
*/
func main() {
	// 客户端
	client := &http.Client {
	}
	url := "http://127.0.0.1:8000/healthz"
	method := "GET"
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("ApiToken", "asda45asdeda54asd8e41d")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
