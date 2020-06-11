package main

import "fmt"

//26.删除排序数组中的重复项

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 0
	for j:=1;j<len(nums);j++  {
		if nums[j-1] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}

func removeDuplicates1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 0
	for j:=1;j<len(nums);j++  {
		if nums[i] != nums[j] {
			i ++
			nums[i] = nums[j]
		}
	}
	return i+1
}

func main()  {
	a := []int{0,0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates1(a))
	fmt.Println(a)
}