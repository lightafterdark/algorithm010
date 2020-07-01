package main

import "fmt"

//1. 两数之和
func twoSum(nums []int, target int) []int {
	res := make([]int,0)
	data := make(map[int]int)
	for k,v := range nums {
		if index,ok := data[target-v];ok&&index!=k{
			res = append(res,index)
			res = append(res,k)
		}
		data[v] = k
	}
	return res
}

func main()  {
	fmt.Println(twoSum([]int{2,7,11,15},9))
}