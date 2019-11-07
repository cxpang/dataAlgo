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
	sort(a, n)
	fmt.Println("spend time", time.Since(begin))
}
func sort(a []int, n int) {
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}
