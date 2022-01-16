package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
Go语言通过首字母的大小写来控制访问权限。无论是方法，变量，常量或是自定义的变量类型，如果首字母大写，则可以被外部包访问，反之则不可以。
而结构体中的字段名，如果首字母小写的话，则该字段无法被外部包访问和解析
最后，如果没有特别的访问控制的话，建议字段名首字母都使用大写字母，从而避免无法解析的错误
 */

func GetMessageFromServer() {
	//resp, _ := http.Get("http://www.baidu.com")
	//fmt.Println(resp)
	apiUrl := "http://127.0.0.1:8000/go"
	// 200 OK
	params := url.Values{}
	params.Add("userId", "00512765")
	params.Add("token", "00512765_token")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	u.RawQuery = params.Encode()

	fmt.Println(u.String())

	clientStudent := Student{
		Id:         1001,
		Name:       "张三",
		Age:        12,
		DataSource: "client",
	}
	b, err := json.Marshal(clientStudent)
	if err != nil {
		fmt.Println("to json failed, err:%v\n", err)
		return
	}
	fmt.Println("client send message ", b)

	contentType := "application/json;charset=utf-8"
	body := bytes.NewBuffer(b)
	resp, err := http.Post(u.String(), contentType, body)

	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}

	defer resp.Body.Close()
	// 接收服务端信息
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		var myStudent Student
		err:= json.Unmarshal(b, &myStudent)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:%v\n", err)
			return
		}
		fmt.Println("client receive student: ", myStudent)
	}
}
