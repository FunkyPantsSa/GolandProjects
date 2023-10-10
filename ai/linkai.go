package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func linkai() (ouput string) {

	url := "https://api.link-ai.chat/v1/chat/completions"

	payload := strings.NewReader("{\n  \"app_code\": \"default\",\n  \"messages\": [\n    {\n      \"role\": \"user\",\n      \"content\": \"你好\"\n    }\n  ]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer Link_RhEfGRxJi9Tvq1XKbEV8c2yvil4nbaDpwhII4FxecE")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	//fmt.Println(string(body))
	return string(body)
}

func main() {
	//var txt string
	//fmt.Print("请输入一个整数：")
	//_, err := fmt.Scan(&txt)
	//if err != nil {
	//	fmt.Println("输入错误：", err)
	//	return
	//}
	println(linkai())
}
