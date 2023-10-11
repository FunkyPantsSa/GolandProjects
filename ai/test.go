package main

import (
	"encoding/json"
	"fmt"
)

type MyType2 struct {
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func main() {
	jsonStr := `{
		"choices":[
			{
				"index":0,
				"message":{
					"role":"assistant",
					"content":"你好！有什么我可以帮助你的吗？"
				},
				"finish_reason":"stop"
			}
		],
		"usage":{
			"prompt_tokens":13,
			"completion_tokens":18,
			"total_tokens":31
		}
	}`
	var data MyType2
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	content := data.Choices[0].Message.Content
	fmt.Println("Content:", content)

}
