package createTicket

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateTicket() [5]int {
	//随机生成数组
	var numbers [5]int
	for i := 0; i <= 4; i++ {
		rand.Seed(time.Now().UnixMilli())
		b := rand.Intn(35)
		numbers[i] = b
		time.Sleep(1 * time.Second)
	}
	return numbers
}

func CreateTicketMore() map[int]any {
	var ticket = make(map[int]any)
	for i := 0; ; i++ {
		var machineTicket []int
		//var numbers [5]int
		for j := 0; j <= 4; j++ {
			rand.Seed(time.Now().UnixNano())
			b := rand.Intn(35)
			//machineTicket[j] = b
			machineTicket = append(machineTicket, b)
			time.Sleep(12 * time.Nanosecond)

		}

		//用map来接收
		ticket[i] = []any{machineTicket}
		//fmt.Println("机选为", ticket)
		//fmt.Println(ticket[i])
	LOOP:
		fmt.Println("是否继续机选？1.是 2.否")
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
