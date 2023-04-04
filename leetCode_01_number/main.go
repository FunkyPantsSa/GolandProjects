package main

import (
	"fmt"
)

var number = []int{2, 4, 11, 3}

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
		//println(i)
		for j := 0; j < len(nums)-1; j++ {
			if j == i {
				j++
			}
			//println(nums[i], nums[j])
			if nums[i]+nums[j] == target {
				//result :=[i,j]
				d = j
				//println(j)
				c = i
				k = []int{c, d}
				break

			}
			//log.Printf("没有发现")

		}

	}
	//k = []int{c, d}
	return k

}
