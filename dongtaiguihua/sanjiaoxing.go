package main

import "math"

func maxProduct(nums []int) int {
	max,imax,imin:=nums[0],nums[0],nums[0]
	for i:=0;i<len(nums);i++{
		if nums[i] < 0 {
			imax,imin = imin,imax
		}
		imax=max_num(imax*nums[i],nums[i])
		imin=min_num(imin*nums[i],nums[i])
		max=max_num(max,imax)
	}
	return max

}

func max_num(x,y int)int {
	max:=x
	if x<y {
		max = y
	}
	return max
}
func min_num(x,y int)int{
	min:=x
	if x<y {
		min = y
	}
	return y
}


/****

class Solution {
    public int maxProduct(int[] nums) {
        int max = Integer.MIN_VALUE, imax = 1, imin = 1;
        for(int i=0; i<nums.length; i++){
            if(nums[i] < 0){
              int tmp = imax;
              imax = imin;
              imin = tmp;
            }
            imax = Math.max(imax*nums[i], nums[i]);
            imin = Math.min(imin*nums[i], nums[i]);

            max = Math.max(max, imax);
        }
        return max;
    }
}

作者：guanpengchn
链接：https://leetcode-cn.com/problems/maximum-product-subarray/solution/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */