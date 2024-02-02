package contrastTicket

import (
	"fmt"
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

func ConTrastMore(a map[int]any, b [5]int) (bool, [5]int) {
	var ticket [5]int
	var result bool
	for i := 0; i <= len(a); i++ {
		c := a[i]
		fmt.Println(c)

		for j := 0; j < 5; j++ {
			//fmt.Printf("你选择的号码第一位是%d \n", c)
			//fmt.Printf("机器选择的号码第一位是%d \n", b[i])
			if a[i] == b[i] {
				result = true
			} else {
				result = false
				break // 只要有一个不相等，就跳出循环
			}

		}

	}
	return result, ticket

}
