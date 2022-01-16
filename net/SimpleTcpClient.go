package net

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ConnectToServer(endpoint string) {
	con, err := net.Dial("tcp", endpoint)
	if err != nil {
		fmt.Println("connect to %s failed, error %s", endpoint, err)
	}
	defer con.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = con.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := con.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
