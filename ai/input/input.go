package input

import (
	"bufio"
	"fmt"
	"os"
)

func Input() (input string) {
	// 创建一个 bufio.Scanner 对象，用于从标准输入读取用户输入
	reader := bufio.NewScanner(os.Stdin)
	//fmt.Printf("我：")
	// 循环读取用户输入
	for reader.Scan() {
		// 获取用户输入的文本
		input := reader.Text()
		return input
		// 处理用户输入
		//fmt.Printf("你输入了: %s\n", input)
	}
	// 检查是否读取到了用户输入
	if err := reader.Err(); err != nil {
		fmt.Printf("读取用户输入出错: %v\n", err)
	}
	return

}

func InputInt() (input []byte) {
	// 创建一个 bufio.Scanner 对象，用于从标准输入读取用户输入
	reader := bufio.NewScanner(os.Stdin)
	//fmt.Printf("我：")
	// 循环读取用户输入
	for reader.Scan() {
		// 获取用户输入的文本
		input := reader.Bytes()
		return input
		// 处理用户输入
		//fmt.Printf("你输入了: %s\n", input)
	}
	// 检查是否读取到了用户输入
	if err := reader.Err(); err != nil {
		fmt.Printf("读取用户输入出错: %v\n", err)
	}
	return

}
