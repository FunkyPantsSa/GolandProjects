package main

import (
	"fmt"
	"github.com/yalp/jsonpath/api2d"
	"github.com/yalp/jsonpath/input"
)

func main() {
	fmt.Println("欢迎使用AI小助手")
	//input.Input()
	fmt.Println("请输入您的token")
	var token = input.Input()
	for {

		fmt.Printf("我：")
		fmt.Println("AI:", api2d.Api2d(token, input.Input()))

	}
}
