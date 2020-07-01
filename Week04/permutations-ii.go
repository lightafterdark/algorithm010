package main

import (
	"fmt"
	"sort"
)

//47. 全排列 II

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}else if n == 1 {
		return [][]int{nums}
	}
	res := make([][]int,0)
	sort.Ints(nums)

	//i表示本次函数需要放置得元素位置
	var helper  func(i int)
	helper = func(i int) {
		if i == n - 1 {
			tmp := make([]int,n)
			copy(tmp,nums)
			res = append(res,tmp)
		}

		//nums[0:i]是已经决定得部分,nums[i:]是待决定部分,同时待选元素也都在nums[i:]
		for k:=i;k<n ;k++  {
			//跳过重复数字
			if k != i && nums[k] == nums[i] {
				continue
			}
			nums[k],nums[i] = nums[i],nums[k]
			helper(i+1)
		}
		//还原状态
		for k := n-1;k>i;k-- {
			nums[i],nums[k] = nums[k],nums[i]
		}
	}
	helper(0)
	return res
}

func main()  {
	fmt.Println(permuteUnique([]int{1,1,2}))
}

