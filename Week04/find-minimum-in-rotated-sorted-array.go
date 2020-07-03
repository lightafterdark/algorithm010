package main

import "fmt"

//153. 寻找旋转排序数组中的最小值

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums)-1
	if nums[right] > nums[0] {
		return nums[0]
	}

	for left < right {
		mid := left + (right - left)/2
		if nums[mid] > nums[mid + 1 ]{
			return nums[mid+1]
		}

		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[left]{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return -1
}



func main()  {
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
}