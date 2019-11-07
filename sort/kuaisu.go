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
	QuickSort(a, n)
	fmt.Println(a)
	fmt.Println("spend time", time.Since(begin))
}

func QuickSort(a []int,n int)  {
	QuickSortC(a,0,n-1)
}
func QuickSortC(a []int,p,r int)  {
	if p>r {
		return
	}
	q := partition(a, p, r) // 获取分区点 quick_sort_c(A, p, q-1) quick_sort_c(A, q+1, r)
	QuickSortC(a,p,q-1)
	QuickSortC(a,q+1,r)
}
func partition(a []int,l,r int) int {
    value:=a[r]
    i:=l
	for j:=l;j<r ;j++  {
		if a[j]<value{
			a[i],a[j] = a[j],a[i]
			i++
		}
	}
	a[i],a[r]=a[r],a[i]
	return i

}
