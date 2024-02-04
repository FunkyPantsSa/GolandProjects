package contrastTicket

func PrizeLevel(a []int, b [5]int) string {
	var prizeLevel string
	var j int = 0
	// 假设这里有一些逻辑来决定奖级
	for i := 0; i < 5; i++ {

		if a[i] == b[i] {
			j++
		}

	}
	if j < 3 {

		prizeLevel = "0"
	} else if j == 3 {

		prizeLevel = "3"
	} else if j == 4 {

		prizeLevel = "2"
	} else if j == 5 {

		prizeLevel = "1"
	}

	return prizeLevel
}
