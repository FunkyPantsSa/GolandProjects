package main

import (
	"fmt"
	"time"
)

func range1() {
	var m = []int{1, 2, 3, 4, 5}

	{
		i, v := 0, 0
		for i, v = range m {
			go func() {
				time.Sleep(time.Second * 3)
				fmt.Println(i, v)
			}()
		}
	}

	time.Sleep(time.Second * 10)
}

func range2() {
	var m = []int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(time.Second * 10)
}

func main() {
	range2()
	//4 5
	//4 5
	//4 5
	//4 5
	//4 5

	//range1()
	//2 3
	//0 1
	//1 2
	//3 4
	//4 5
}
