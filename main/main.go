package main

import (
	"basic/context"
	"basic/net"
)

func main() {
	context.TestBasicContext()
}

// 测试简单的TCP功能
func testTcp() {
	endPoint := "127.0.0.1:8080"
	go net.InitServer(endPoint)
	net.ConnectToServer(endPoint)
}

// 测试简单的UDP功能
func testUdp() {
	go net.InitUdpServer()
	net.ConnectToUdpServer()
}

/*
Hello server 127.0.0.1:54115 12
recv:Hello server addr:127.0.0.1:8081 count:12
*/

// 测试简单的HTTP功能
func testHttp() {
	go net.InitHttpServer()
	net.GetMessageFromServer()
}

/*
127.0.0.1:58214 连接成功
method: GET
url: /go
header: map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]]
body: {}
200 OK
map[Content-Length:[12] Content-Type:[text/plain; charset=utf-8] Date:[Sun, 09 J
an 2022 13:45:16 GMT]]
读取完毕
www.5lmh.com
*/
