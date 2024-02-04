package contrastTicket

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Contrast(a [5]int, b [5]int) bool {
	//对比数组是否一致
	var result bool
	for i := 4; i >= 0; i-- {
		if a[i] == b[i] {
			result = true
		} else {
			result = false
			break // 只要有一个不相等，就跳出循环
		}

	}
	return result

}

func OneContrast(a [5]int, b [5]int) bool {
	//对比数组是否一致
	var result bool
	for i := 0; i < 5; i++ {

		fmt.Printf("你选择的号码第一位是%d \n", a[i])
		fmt.Printf("机器选择的号码第一位是%d \n", b[i])
		if a[i] == b[i] {
			result = true
		} else {
			result = false
			break // 只要有一个不相等，就跳出循环
		}

	}
	return result

}

// 对比多张
func ConTrastMore(a map[int]any, b [5]int) (bool, []int) {
	var ticket []int
	var result bool
	var k int
	//fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		c := CutManyTicket(MapToArr(a[i]))
		//fmt.Println("c=", c)
		for j := 0; j < 5; j++ {

			if c[i] == b[j] {
				k++
				fmt.Println(c[i], b[j])

			}

			//fmt.Printf("你选择的号码第一位是%d \n", c[j])
			//fmt.Printf("机器选择的号码第一位是%d \n", b[j])
			//if c[i] == b[i] {
			//
			//	result = true
			//	ticket = c
			//	fmt.Println("您选择的彩票", c, "中奖啦！！！！！")
			//} else {
			//	result = false
			//	fmt.Println("您选择的彩票", c, "没有中奖")
			//	ticket = nil
			//	break // 只要有一个不相等，就跳出循环
			//}

		}
		if k > 1 {
			result = true
			ticket = c
		} else {
			result = false
			ticket = nil
		}
	}

	return result, ticket
}

// map更改为数组
func MapToArr(a interface{}) string {
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

// 分割，然后加进数组
func CutManyTicket(a string) []int {
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
			return nil
			// 或者你可以选择忽略错误，继续处理下一个数字
		}
		// 将转换后的整数添加到整数切片中
		numbers = append(numbers, num)
	}

	// 打印提取的数字
	//fmt.Println(numbers)

	return numbers
}
