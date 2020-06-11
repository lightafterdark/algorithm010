package main

import "fmt"

//70爬楼梯




//递归，会超时
func climbStairs1(n int) int {
	return helper1(0,n)
}

func helper1(i,n int)  int{
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return helper1(i+1,n)+helper1(i+2,n)
}


//记忆姓递归
func climbStairs2(n int) int {
	memo := make(map[int]int)
	return helper2(0,n,memo)
}

func helper2(i,n int,memo map[int]int)  int{
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if _,ok := memo[i];ok {
		return memo[i]
	}
	memo[i] = helper2(i+1,n,memo)+helper2(i+2,n,memo)
	return memo[i]
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}else{
		one := 1
		two := 2
		res := 0
		for i:=2;i<n ;i++  {
			res = one + two
			one,two = two,res
		}
		return res
	}
}




func main()  {
	//fmt.Println(climbStairs2(44))
	fmt.Println(climbStairs(100))
}

