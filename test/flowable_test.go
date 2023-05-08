package test

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestFlowable(t *testing.T) {
	// Flowable REST API的URL
	url := "http://localhost:9000/flowable-rest/service/identity/users"

	// 设置HTTP Basic Authentication标头
	username := "rest-admin"
	password := "test"
	auth := username + ":" + password
	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	authHeader := "Basic " + base64Auth

	// 创建HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(authHeader)
	// 添加Authorization标头
	req.Header.Add("Authorization", authHeader)

	// 发送HTTP请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 处理HTTP响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
