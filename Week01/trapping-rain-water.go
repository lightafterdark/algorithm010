package main

import (
	"fmt"
)

//42.接雨水
func max(x,y int)  int{
	if x > y {
		return x
	}else{
		return y
	}
}
func trap(height []int) int {
	res := 0
	max_left := 0
	max_right := 0
	left := 1
	right := len(height)-2

	for i:=1;i<len(height)-1;i++ {
		if(height[left-1]<height[right+1]){
			max_left = max(max_left,height[left-1])
			min := max_left
			if min > height[left]{
				res += (min-height[left])
			}
			left++
		}else{
			max_right = max(max_right,height[right+1])
			min := max_right
			if min > height[right] {
				res += (min-height[right])
			}
			right--
		}
	}
	return res
}

func main()  {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
