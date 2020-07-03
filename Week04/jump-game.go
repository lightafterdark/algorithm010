package main

import "fmt"

//55. 跳跃游戏

func canJump(nums []int) bool {
	n := len(nums)
	rightMost := 0
	for k,v := range nums {
		if k <= rightMost {
			if rightMost > k+v {
				rightMost = rightMost
			}else{
				rightMost = k + v
			}
			if rightMost >= n - 1 {
				return true
			}
		}
	}
	return false
}	


func main()  {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
	fmt.Println(canJump([]int{2,0}))
	fmt.Println(canJump([]int{2,5,0,0}))
}