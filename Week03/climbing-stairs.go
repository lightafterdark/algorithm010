package main

import "fmt"

//70. 爬楼梯

//常规递归，超时
//func climbStairs(n int) int {
//	res := 0
//	if n <= 0 {
//		return 0
//	}else if n == 1 {
//		return 1
//	}else if n == 2 {
//		return 2
//	}
//	res = climbStairs(n-1)+climbStairs((n-2))
//	return res
//}

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	mem := make(map[int]int)
	mem[1]=1
	mem[2]=2
	return helper(n,mem)
}

func helper(n int,mem map[int]int)  int{
	if num,ok := mem[n];ok{
		return num
	}
	mem[n] = helper(n-1,mem)+helper(n-2,mem)
	return mem[n]
}

func main()  {
	fmt.Println(climbStairs(38))
}