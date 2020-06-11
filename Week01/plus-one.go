package main

import (
	"fmt"
)

//66. 加一

func plusOne(digits []int) []int {
	length := len(digits)
	index := length -1
	for true {
		if index == 0 &&  digits[0] == 9 {
			digits[0] = 0
			res := make([]int,0)
			res = append(res,1)
			res = append(res,digits...)
			return res
		}else if digits[index] == 9 {
			digits[index] = 0
			index--
		}else{
			digits[index] = digits[index] + 1
			break
		}
	}
	return digits
}

func main()  {
	a := []int{9}
	fmt.Println(plusOne(a))
}