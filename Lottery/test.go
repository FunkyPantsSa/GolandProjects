package main

import (
	"Lottery/createTicket"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func generate() (interface{}, bool) {
	s := []string{"123", "345", "abc"}
	//s := 123
	//s := "mmm"
	return s, true
}

func test(a interface{}) string {
	origin, ok := a, true
	var k reflect.Value
	var m string
	if ok {

		switch reflect.TypeOf(origin).Kind() {
		case reflect.Slice, reflect.Array:
			s := reflect.ValueOf(origin)
			for i := 0; i < s.Len(); i++ {

				//fmt.Println(s.Index(i))
				k = s.Index(i)
				k2 := fmt.Sprintf("%v", k)
				//fmt.Println(k2)
				return k2

			}
		case reflect.String:
			s := reflect.ValueOf(origin)
			fmt.Println(s.String(), "I am a string type variable.")
		case reflect.Int:
			s := reflect.ValueOf(origin)
			t := s.Int()
			fmt.Println(t, " I am a int type variable.")
		}
	}
	return m
}

func CutManyTicket(a string) ([]int, error) {
	// 假设你有一个数字字符串
	numbersStr := a

	newStr := strings.Replace(numbersStr, "[", "", -1)
	newStr = strings.Replace(newStr, "]", "", -1)
	// 使用 strings.Split 分割字符串
	//fmt.Println(newStr)
	numbersStrSlice := strings.Split(newStr, " ")

	// 创建一个整数切片来存储提取的数字
	var numbers []int

	// 遍历字符串切片，并将每个字符串转换为整数
	for _, numStr := range numbersStrSlice {
		// 使用 strconv.Atoi 转换字符串为整数
		num, err := strconv.Atoi(numStr)
		if err != nil {
			// 如果转换失败，打印错误并退出
			fmt.Println("Error converting string to int:", err)
			return nil, err
			// 或者你可以选择忽略错误，继续处理下一个数字
		}
		// 将转换后的整数添加到整数切片中
		numbers = append(numbers, num)
	}

	// 打印提取的数字
	fmt.Println(numbers)

	return numbers, nil
}

func main() {
	a := createTicket.CreateTicketMore()
	//b := createTicket.CreateTicket()
	//
	//fmt.Println(a, b)
	//fmt.Println(contrastTicket.ConTrastMore(a, b))
	for i := 0; i < len(a); i++ {

		ticket, err := CutManyTicket(test(a[i]))
		if err != nil {
			return
		}
		fmt.Println(ticket)

	}
	//CutManyTicket(test(a[1]))
}
