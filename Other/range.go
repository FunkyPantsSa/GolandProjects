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
			fmt.Println(i, v)
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
	//range2()
	//4 5
	//4 5
	//4 5
	//4 5
	//4 5

	range1()
	//2 3
	//0 1
	//1 2
	//3 4
	//4 5
}

//https://studygolang.com/articles/28128
//这2段代码实际上是遍历数组的所有变量。
//由于闭包只是绑定到这个value变量上，并没有被保存到goroutine栈中， 所以以上代码极有可能运行的结构都输出为切片的最后一个元素。
//因为这样写会导致for循环结束后才执行goroutine多线程操作，这时候value值只指向了最后一个元素。
//这样的结果不是我们所希望的，而且还会产生并发的资源抢占冲突所以是非常不推荐这样写的。
