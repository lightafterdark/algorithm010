package main

import "fmt"

//239.滑动窗口最大值

//暴力法
func max(x,y int)  int{
	if x > y {
		return x
	}else{
		return y
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	n := len(nums)
	if n*k == 0 {
		return res
	}

	for i:=0;i<n-k+1 ;  i++{
		max := max()
	}
	return res
}


//双向队列


//堆（优先队列）

func main()  {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3))
}