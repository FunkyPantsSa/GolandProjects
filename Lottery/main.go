package main

import (
	"Lottery/buyTicket"
	"Lottery/createTicket"
	"fmt"
)

//func CreateTicket() [5]int {
//	var numbers [5]int
//	for i := 0; i <= 4; i++ {
//
//		b := rand.Intn(35)
//		numbers[i] = b
//	}
//	return numbers
//}

func printArr(arrInput [5]int) {
	//打印数组
	//var arrOutput int
	//println("arrInput:", arrInput)
	for _, arrOutput := range arrInput {
		//println("  ")
		fmt.Println(arrOutput)
		//print(arrOutput)

	}

}

func main() {
	//var a, b [5]int
	fmt.Println("--------------------")
	//printArr(CreateTicket())
	//fmt.Println(createTicket.CreateTicket())
	//fmt.Println(buyTicket.SelfChoice())
	//a = [5]int{1, 2, 3, 2, 5}
	//b = [5]int{1, 2, 3, 4, 5}
	//fmt.Println(buyTicket.SelfChoice2())
	//buyTicket.SelfChoice2()
	fmt.Println("欢迎购彩")
	var a int
LOOP:
	fmt.Println("选择购彩方式1.机选 2.自选")
	fmt.Scanf("%d", &a)
	if a == 2 {
		v := buyTicket.SelfChoice2()
		//fmt.Println(len(v))
		fmt.Println(v[1])
	} else if a == 1 {
		k := createTicket.CreateTicket()
		fmt.Println(k)
	} else {
		fmt.Println("输入错误，请重新输入。")
		goto LOOP
	}

}
