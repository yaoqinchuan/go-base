package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InitHttpServer() {
	//http://127.0.0.1:8000/go
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

type Student struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	DataSource string `json:"data_source"`
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	fmt.Println("request params is  ", params)
	fmt.Println(r.RemoteAddr, " 连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method: ", r.Method)
	// /go
	fmt.Println("url: ", r.URL.Path)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err:%v\n", err)
		return
	}
	serverStudent := Student{}
	err = json.Unmarshal(b, &serverStudent)
	if err != nil {
		fmt.Println("json err:", err)
	}
	serverStudent.DataSource = "server"
	response, err := json.Marshal(serverStudent)
	if err != nil {
		fmt.Println("json err:", err)
	}
	json.Unmarshal(response, &serverStudent)
	fmt.Println("response ", serverStudent)
	// 回复
	w.Write(response)
}
