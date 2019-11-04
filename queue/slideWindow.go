package main

func main() {
	
}


func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	if len(nums) == 0 {
		return result
	}

	window := []int{}
	for i, x := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) != 0 && nums[window[len(window)-1]] <= x {
			window = []int{}
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result

}
//我的方法
func test(nums []int, k int) []int  {
	var window,res []int
	for index,value:=range nums{
		if index>=k && window[0] <=index-k{   //下标index ，k为窗口值
			window = window[1:]
		}

		for len(window) > 0&&nums[window[len(window)-1]]<=value {
			window = window[0:len(window)-1]
		}
		window = append(window, index) //window窗口存放的是下标
		if index>=k-1{
			res = append(res, nums[window[0]])
		}
	}
	return  res
}