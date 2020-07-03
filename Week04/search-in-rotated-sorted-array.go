package main

import "fmt"

//33. 搜索旋转排序数组
func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	left := 0
	right := len(nums)-1
	for left <= right  {
		middle := (left + right)/2
		if nums[middle] == target {
			return middle
		}
		if nums[0] <= nums[middle] {
			if nums[0] <= target && target <= nums[middle]{
				//目标值在前半段
				right = middle - 1
			}else{
				//目标值在后半段
				left = middle + 1
			}
		}else{
			//说明有变化的前半段
			if nums[middle] < target && target <= nums[len(nums)-1]{
				left = middle + 1
			}else{
				right = middle - 1
			}
		}
	}
	return -1
}

func main()  {
	fmt.Println(search([]int{4,5,6,0,1,2},5))

}