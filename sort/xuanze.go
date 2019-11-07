package main

import (
	"fmt"
	"time"
)

func main() {
	var a []int = []int{3, 2, 5, 6, 4, 8, 7}
	begin := time.Now()
	fmt.Println("before sort")
	fmt.Println(a)
	n := len(a)
	fmt.Println("after sort")
	chsort(a, n)
	fmt.Println("spend time", time.Since(begin))
}
func chsort(values []int, n int) {
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		min := i // 初始的最小值位置从0开始，依次向右
		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := n - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		//fmt.Printf("i:%d min:%d\n", i, min)
		// 把每次找出来的最小值与之前的最小值做交换
		values[i], values[min] = values[min], values[i]
	}
	fmt.Println(values)
}
