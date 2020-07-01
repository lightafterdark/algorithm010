package main

import "fmt"

//46. 全排列

func permute(nums []int) [][]int {
	//如果传入数组长度未一，直接返回
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := make([][]int,0)
	//遍历数组
	for k,v := range nums {
		//每次遍历，当前元素未开始，求剩下元素得排列
		tmp := make([]int,len(nums)-1)
		//copy复制为值复制，新切片和原始切片，互不影响
		copy(tmp[0:],nums[0:k])
		copy(tmp[k:],nums[k+1:])
		sub := permute(tmp)
		for _,s := range sub {
			//合并结果
			res = append(res,append(s,v))
		}
	}
	return res
}

func main()  {
	fmt.Println(permute([]int{1,2,3}))
	
}