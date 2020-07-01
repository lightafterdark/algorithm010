package main

import "fmt"

//22. 括号生成

func generateParenthesis(n int) []string {
	res := make([]string,0)
	var gen func(left,right,n int,s string)
	gen = func(left, right, n int, s string) {
		if left == n && right == n {
			res = append(res,s)
		}

		if left < n {
			gen(left+1,right,n,s+"(")
		}

		if left > right && right < n {
			gen(left,right+1,n,s+")")
		}
	}
	gen(0,0,n,"")
	return res
}

func main()  {
	fmt.Println(generateParenthesis(3))
}