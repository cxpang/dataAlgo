package main

import (
	"fmt"
	"sort"
)

//四个数相加只和
func main() {
	var nums  =  []int{-1,0,1,2,-1,-4}
	res := threeSum(nums)
	fmt.Println(res)
}

//大神算法

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	var(
		i int
		r int
		l int
	)
	var length int = len(nums)

	for i = 0; i < length; i++ {
		if (nums[i]>0) {
			continue
		}
		if (i>0 && nums[i] == nums[i-1]) {
			continue
		}
		l = length-1;
		for r = i+1; r < l;{
			if (l<length-1 && nums[l] == nums[l+1]) {
				l--
				continue
			}
			if (r>i+1 && nums[r] == nums[r-1]) {
				r++
				continue
			}
			sum := nums[i]+nums[r]+nums[l]
			if (sum > 0) {
				l--
			}
			if (sum < 0) {
				r++
			}
			if (sum == 0) {
				arr := []int{nums[i],nums[r],nums[l]}
				res = append(res,arr)
				l--;r++;
			}
		}
	}
	return res
}
//本人解法
func threeSum2(nums []int) [][]int {
	//先对数组排序
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		z := len(nums) - 1
		for z > j {
			b := nums[j]
			c := nums[z]
			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)
				for j < z && nums[j] == nums[j+1] {
					j++
				}
				for j < z && nums[z] == nums[z-1] {
					z--
				}
				j++
				z--
			}
		}
	}
	return result
}