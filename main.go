package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Completion struct {
	Response  string `json:"response"`
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
}

type ChatRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func main() {
	url := "http://localhost:11434/api/generate"
	// 创建请求体结构体
	requestBody := ChatRequest{
		Model:  "deepseek-r1:7b",
		Prompt: "如何做好家里的卫生工作",
		Stream: false,
	}

	// 将结构体转换为 JSON
	payload, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer sk-b8ebb99508964850b2b1c")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var completion Completion
	err = json.Unmarshal(body, &completion)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Content:", completion.Response)
}
