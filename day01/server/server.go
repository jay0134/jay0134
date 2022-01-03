package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回 200
*/

func main() {
	host := flag.String("host", "127.0.0.1", "listen host")
	port := flag.String("port", "8000", "listen port")
	flag.Parse()
	http.HandleFunc("/healthz", Healthz)

	err := http.ListenAndServe(*host+":"+*port, nil)

	if err != nil {
		panic(err)
	}
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 1. 获取请求数据并打印
	fmt.Printf("header:%v \n", r.Header)
	fmt.Printf("host:%v \n", r.Host)
	fmt.Printf("request-url:%v \n", r.RequestURI)
	fmt.Printf("method:%v \n", r.Method)
	//version := os.Environ()
	//for i:= range version{
	//	fmt.Println(version[i])
	//}
	goPath := os.Getenv("GOPATH")
	token := r.Header.Get("api-token")
	requestUrl := r.RequestURI
	host := r.Host
	method := r.Method
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}

	fmt.Println(string(b))
	answer := `{"status": 200}`
	w.Header().Set("api-token", token)
	w.Header().Set("host", host)
	w.Header().Set("Request URL", requestUrl)
	w.Header().Set("Method", method)
	w.Header().Set("GOPATH", goPath)

	w.Write([]byte(answer))
}
