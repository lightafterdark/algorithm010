package main

import "fmt"

//69. x 的平方根


//二分查找
func mySqrt(x int) int {
	left := 0
	right := x
	res := -1
	for left <= right {
		middle := left + (right-left)/2
		if middle*middle <= x {
			//向右找
			res = middle
			left = middle + 1
		}else{
			//像左找
			right = middle - 1
		}
	}
	return res
}

func main()  {
	fmt.Println(mySqrt(8))
}