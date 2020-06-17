package main
//1.两数之和

import "fmt"
//1.两数之和

func twoSum(nums []int, target int) []int {
	res := make([]int,0)
	temp := make(map[int]int)
	for i,num := range nums {
		if index,ok := temp[target-num];ok && i != index {
			res = append(res,i)
			res = append(res,index)
		}
		temp[num] = i
	}
	return res
}

func main()  {
	a := []int{2,7,11,15}
	fmt.Println(twoSum(a,9))
	fmt.Println(twoSum(a,9))

}


