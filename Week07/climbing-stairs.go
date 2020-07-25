package main

import "fmt"

//70. 爬楼梯

//递归
//func climbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//	return climbStairs(n-1) + climbStairs(n-2)
//}

//hashmap
//func climbStairs(n int) int {
//	m := make(map[int]int)
//	m[1] = 1
//	m[2] = 2
//	var helper func(n int) int
//	helper = func(n int) int {
//		if num, ok := m[n]; ok {
//			return num
//		}
//		m[n] = helper(n-1) + helper(n-2)
//		return m[n]
//	}
//	return helper(n)
//}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	left := 1
	right := 2
	res := 0
	for i := 3; i <= n; i++ {
		res = left + right
		left, right = right, res
	}
	return res
}

func main() {
	fmt.Println(climbStairs(44))
}
