package main

import (
	"Lottery/buyTicket"
	"Lottery/contrastTicket"
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
	var userTicket map[int]any
	fmt.Println("--------------------")
	//printArr(CreateTicket())
LOOP:
	fmt.Println("选择购彩方式1.机选 2.自选")
	fmt.Scanf("%d", &a)
	if a == 2 {
		//自己输入想要的数字
		userTicketBySelf := buyTicket.SelfChoice2()
		userTicket = userTicketBySelf
		//fmt.Println(len(v))
		//fmt.Println(userTicketBySelf)
	} else if a == 1 {
		//由机器随机生成数字
		userTicketByMachine := createTicket.CreateTicketMore()
		//fmt.Println(userTicketByMachine)
		userTicket = userTicketByMachine
	} else {
		fmt.Println("输入错误，请重新输入。")
		goto LOOP
	}

	fmt.Println("--------------------")
	//printArr(CreateTicket())
	//fmt.Println("选择是:", userTicket)
	for i := 0; i < len(userTicket); i++ {
		fmt.Println("您的彩票是")
		fmt.Println(contrastTicket.CutManyTicket(contrastTicket.MapToArr(userTicket[i])))
		//contrastTicket.MapToArr(userTicket)
		//printArr(userTicket[i])

	}

	fmt.Println("--------------------")
	//printArr(CreateTicket())
	//time.Sleep(time.Second * 5)
	fmt.Println("开奖啦！！！！！")
	fmt.Println("--------------------")

	vtory := createTicket.CreateTicket()
	fmt.Println("中奖号码是：", vtory)

	fmt.Println("--------------------")
	_, m := contrastTicket.ConTrastMore(userTicket, vtory)
	//for i := 0; i < len(userTicket); i++ {
	if m != nil {
		fmt.Println("恭喜你中奖了，中奖号码是：", m)
		contrastTicket.PrizeLevel(m, vtory)
	} else {
		fmt.Println("很遗憾你没有中奖。")
		fmt.Println("--------------------")

	}

	//fmt.Println(contrastTicket.CutManyTicket(contrastTicket.MapToArr(userTicket[i])))
	//contrastTicket.MapToArr(userTicket)
	//printArr(userTicket[i])

	//}
}
