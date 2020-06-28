package main

import "fmt"

//46. 全排列

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	for i,num := range nums {
		tmp := make([]int,len(nums)-1)
		copy(tmp[0:],nums[0:i])
		copy(tmp[i:],nums[i+1:])

		sub := permute(tmp)
		for _,s := range sub {
			res = append(res,append(s,num))
		}
	}
	return res
}

func main()  {
	fmt.Println(permute([]int{1,2,3}))
}
