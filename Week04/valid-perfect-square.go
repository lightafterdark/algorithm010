package main

import "fmt"

//367. 有效的完全平方数

func isPerfectSquare(num int) bool {
	left := 0
	right := num
	for left <= right {
		//middle := left + (right - left)/2
		middle := (left + right)/2
		if middle*middle == num {
			return true
		}else if middle*middle < num {
			left = middle + 1
		}else{
			right = middle - 1
		}
	}
	return false
}

func main()  {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}