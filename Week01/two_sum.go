package main
//1.两数之和

import "fmt"
//1.两数之和

func twoSum1(nums []int, target int) []int {
    res := make([]int,0)
    for i,n := range nums {
        for ii,nn := range nums {
            if target == n + nn && i != ii {
                res = append(res,i)
                res = append(res,ii)
                return res
            }
        }
    }
    return res
}

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
	sum := twoSum1(a,9)
	fmt.Println(sum)

}


