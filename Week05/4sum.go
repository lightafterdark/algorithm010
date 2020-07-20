package main

import (
	"fmt"
	"sort"
)


//func fourSum(nums []int, target int) [][]int {
//	sort.Ints(nums)
//	var ret [][]int
//	if len(nums)<4{
//		return ret
//	}
//	for i:=0;i<len(nums)-4 ;i++ {
//		for j:=1;j<len(nums)-3 ;j++  {
//			for m:=2;m<len(nums)-2 ;m++  {
//				for n:=3;n<len(nums)-1 ;n++  {
//					if nums[i]+nums[j]+nums[m]+nums[n] == target {
//						ret = append(ret,[]int{nums[i],nums[j],nums[m],nums[n]})
//					}
//				}
//			}
//		}
//	}
//	return ret
//}

func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	if len(nums)<4{
		return ret
	}
	sort.Ints(nums)
	for i:=0;i<len(nums)-4 ;i++ {
		for j:=i+1;j<len(nums)-3 ;j++  {
			for m:=j+1;m<len(nums)-2 ;m++  {
				for n:=m+1;n<len(nums)-1 ;n++  {
					if nums[i]+nums[j]+nums[m]+nums[n] == target {
						ret = append(ret,[]int{nums[i],nums[j],nums[m],nums[n]})
					}
				}
			}
		}
	}
	return ret
}

func main()  {
	fmt.Println(fourSum([]int{1,0,-1,0,-2,2},0))
}
