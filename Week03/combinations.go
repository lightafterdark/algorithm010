package main

import "fmt"

//77. 组合

func combine(n int, k int) [][]int {
	res := make([][]int,0)
	tmp := make([]int,k)
	var dfs func(start int,index int)
	dfs = func(start int, index int) {
		for i:=start; i <= n-(k-1-index);i++  {
			tmp[index] = i
			if index == k - 1 {
				res = append(res,append([]int{},tmp...))
			}else {
				dfs(i+1,index+1)
			}
		}
	}
	dfs(1,0)
	return res
}


func main()  {
	fmt.Println(combine(4,2))

}