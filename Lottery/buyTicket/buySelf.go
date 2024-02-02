package buyTicket

import (
	//"crypto/rand"
	"fmt"
	"reflect"
)

//自己选择彩票额数据

func SelfChoice() []int {
	var selfTicket []int
	fmt.Println("请输入您选择的彩票号码，共5个数字，数字范围为1-35")
	for i := 0; i <= 4; i++ {
		var t int
		fmt.Scanf("%d", &t)
		//判断输入的t的数据类型和t的取值大小
		if reflect.ValueOf(t).Kind() != reflect.Int || t < 1 || t > 35 {
			fmt.Println("输入错误，请重新输入。")
			i--
			continue
		}
		selfTicket = append(selfTicket, t)
	}
	//fmt.Println("您选择的彩票号码是：", selfTicket)
	return selfTicket

}

type arr []int

func SelfChoice2() map[int]any {
	//map[string]int
	var ticket = make(map[int]any)
	for i := 0; ; i++ {
		var selfTicket2 []int
		fmt.Printf("请输入您选择的%d彩票号码，共5个数字，数字范围为1-35 \n", i)
		for j := 0; ; j++ {
			var t, k int
			fmt.Scanf("%d", &t)
			k = i
			//判断输入的t的数据类型和t的取值大小
			if reflect.ValueOf(t).Kind() != reflect.Int || t < 1 || t > 35 {
				fmt.Println("输入错误，请重新输入。")
				i = k
				continue
			}
			selfTicket2 = append(selfTicket2, t)
			if len(selfTicket2) == 5 {
				break
			}

		}
		//用map来接收
		ticket[i] = []arr{selfTicket2}
		fmt.Println(ticket[i])
	LOOP:
		fmt.Println("是否继续输入？1.是 2.否")
		var a int
		fmt.Scanf("%d", &a)
		if a == 2 {
			break
		} else if a == 1 {

		} else {
			fmt.Println("输入错误，请重新输入。1.是 2.否")
			goto LOOP
		}

	}
	//ticket[1] = 1
	//ticket[2] = 2
	//fmt.Println(ticket)
	//
	return ticket

}
