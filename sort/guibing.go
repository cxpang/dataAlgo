package main

import (
	"fmt"
	"time"
)

func main()  {
	var a []int = []int{3, 2, 5, 6, 4, 8,9, 7}
	begin := time.Now()
	fmt.Println("before sort")
	fmt.Println(a)
	n := len(a)
	fmt.Println("after sort")
	MergeSort(a, 0,n-1)
	fmt.Println(a)
	fmt.Println("spend time", time.Since(begin))
}
func MergeSort(a []int,l ,r int)  {
	if l>=r{
		return
	}
	mid:=(r+l)/2
	MergeSort(a, l, mid)
	MergeSort(a, mid+1, r)
	merge(a,l,mid,r)
}
func merge(arr []int, l int, mid int, r int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	// 指向两部分起点
	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		// 左边的点超过中点，说明只剩右边的数据
		if left > mid {
			arr[i] = temp[right-l]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > r {
			arr[i] = temp[left-l]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left - l] > temp[right - l] {
			arr[i] = temp[right - l]
			right++
		} else {
			arr[i] = temp[left - l]
			left++
		}
	}
}

