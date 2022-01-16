package net

import (
	"bufio"
	"fmt"
	"net"
)

// 这里的是协程的概念，比线程还要轻量化
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		message := string(buf[:n])
		fmt.Println("收到client端发来的数据：", message)
		conn.Write([]byte(message)) // 发送数据
	}
}

func InitServer(endPoint string) {
	listen, err := net.Listen("tcp", endPoint)
	if err != nil {
		fmt.Println("listen failed, error: ", err)
		return
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, error: ", err)
			continue
		}
		go process(con)
	}
}
