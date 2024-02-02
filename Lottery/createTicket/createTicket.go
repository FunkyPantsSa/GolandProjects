package createTicket

import (
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
