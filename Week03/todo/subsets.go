package main

import "fmt"

//78. 子集

func subsets(nums []int) [][]int {
	res := make([][]int,0)
	//n := len(nums)
	for _,num := range nums {
		for _,curr := range res {
			res = append(res,append(curr,num))
		}
	}
	return res
}

func main()  {
	fmt.Println(subsets([]int{1,2,3}))

}