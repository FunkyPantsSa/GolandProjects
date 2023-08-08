package main

import (
	"fmt"
)

var number = []int{2, 6, 11, 4}
var number2 bool = false

func main() {
	//result =
	fmt.Println(twoSum(number, 6))
	//twoSum(number, 6)

}

func twoSum(nums []int, target int) []int {
	var c, d int
	var k = []int{}
	for i := 0; i < len(nums); i++ {
		//arrHaiCoder[i]
		if number2 == true {
			break
		}
		//println(i)
		for j := 0; j < len(nums); j++ {
			if j == i {
				j++
			}
			//println(nums[i], nums[j])
			if nums[i]+nums[j] == target {
				//result :=[i,j]
				d = j
				//println("j=", j)
				c = i
				k = []int{c, d}
				number2 = true

			}
			//log.Printf("没有发现")

		}

	}
	//k = []int{c, d}
	return k

}
