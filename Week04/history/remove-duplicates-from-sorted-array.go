package main

import "fmt"

//26. 删除排序数组中的重复项

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 0
	for j:=1;j<len(nums);j++  {
		if nums[i] != nums[j]{
			//只保留不相等的数字
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}

func main()  {
	fmt.Println(removeDuplicates([]int{1,1,2,3}))
}