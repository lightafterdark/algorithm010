package main

import "fmt"

//704. 二分查找

func search(nums []int, target int) int {
	//if len(nums) == 1 {
	//	if nums[0] == target {
	//		return 0
	//	}else{
	//		return -1
	//	}
	//
	//}
	if len(nums) == 0 {
		return -1

	}
	left := 0
	right := len(nums)-1

	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			left = mid+1
		}else{
			right = mid-1
		}
	}
	return -1
}

func main()  {
	fmt.Println(search([]int{5},6))
}