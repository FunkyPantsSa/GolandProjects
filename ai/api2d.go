package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int     `json:"prompt_tokens"`
		CompletionTokens int     `json:"completion_tokens"`
		TotalTokens      int     `json:"total_tokens"`
		PreTokenCount    int     `json:"pre_token_count"`
		PreTotal         float64 `json:"pre_total"`
		AdjustTotal      float64 `json:"adjust_total"`
		FinalTotal       int     `json:"final_total"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

func api2d() {
	var a string
	a = "你好"
	url := "https://openai.api2d.net/v1/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer fk196722-qZJvA7TXxjAAv0D3SzQYS2h0c4fKLnFa", // 把 fkxxxxx 替换成你自己的 Forward Key，注意前面的 Bearer 要保留，并且和 Key 中间有一个空格。
	}

	data := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": a},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	//fmt.Println("Status Code:", resp.StatusCode)
	//fmt.Println("JSON Response :", responseData)

}

func jsonToData(body string) (data string) {
	//var body string
	//body = `{"id":"chatcmpl-6vbyHZDboUPHcy5jWDtv0ToyiTrmQ","object":"chat.completion","created":1679188677,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":18,"completion_tokens":245,"total_tokens":263,"pre_token_count":4096,"pre_total":41,"adjust_total":38,"final_total":3},"choices":[{"message":{"role":"assistant","content":"\n\n我为您准备了一个笑话，希望您喜欢：\n\n有一天，一只羊和一只鸭子一起去逛街，走着走着突然发现迷路了。羊想要找人问路，可鸭子说：“别找人啊，他们只会说人话，不一定知道动物该怎么走。”羊听了还觉得有道理，于是就问：“那我们该怎么办？”鸭子举了举爪子说：“我们可以问一下天上的云彩，它们肯定知道路该怎么走。” 这时候，一位路人经过看到他们，问道：“你们两个在做什么呢？”羊和鸭子一脸尴尬，最终还是选择“礼貌性地”问路了。"},"finish_reason":"stop","index":0}]}`
	content := MyResponse{}
	if err := json.Unmarshal([]byte(body), &content); err != nil {
		panic(err)
	}

	//data := gjson.Get(content, "Message").Array()[2]
	return data

}

func main() {
	var body string
	body = `{"id":"chatcmpl-6vbyHZDboUPHcy5jWDtv0ToyiTrmQ","object":"chat.completion","created":1679188677,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":18,"completion_tokens":245,"total_tokens":263,"pre_token_count":4096,"pre_total":41,"adjust_total":38,"final_total":3},"choices":[{"message":{"role":"assistant","content":"\n\n我为您准备了一个笑话，希望您喜欢：\n\n有一天，一只羊和一只鸭子一起去逛街，走着走着突然发现迷路了。羊想要找人问路，可鸭子说：“别找人啊，他们只会说人话，不一定知道动物该怎么走。”羊听了还觉得有道理，于是就问：“那我们该怎么办？”鸭子举了举爪子说：“我们可以问一下天上的云彩，它们肯定知道路该怎么走。” 这时候，一位路人经过看到他们，问道：“你们两个在做什么呢？”羊和鸭子一脸尴尬，最终还是选择“礼貌性地”问路了。"},"finish_reason":"stop","index":0}]}`

	jsonToData(body)
}
