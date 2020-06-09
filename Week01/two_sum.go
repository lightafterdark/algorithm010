package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int,0)
	temp := make(map[int]int)
	for i,n := range nums {
		if index,ok := temp[target - n];ok && i != index {
			res = append(res,i)
			res = append(res,index)
		}
		temp[n] = i
	}
	return res
}

func main()  {
	a := []int{2,7,11,15}
	sum := twoSum(a,9)
	fmt.Println(sum)

}


