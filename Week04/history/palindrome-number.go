package main

import "fmt"

//9. 回文数

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0 ){
		return false
	}
	rx := 0
	for x > rx {
		rx = rx * 10 + x % 10
		x /= 10
	}
	fmt.Println(rx)
	return x == rx || x == rx/10
}

func main()  {
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(1221))
}