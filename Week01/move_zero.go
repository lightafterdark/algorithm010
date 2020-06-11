package main

import "fmt"

//283移动零


/**
从左往右遍历数组，如果是0则不处理
非0则需要向后移动，但是移动到哪里的，所以需要一个下标，index就是记录遍历过的0
 */
func moveZeroes1(nums []int)  {
	index := int(0)
	for i,_ := range nums {
		if nums[i] == 0 {
			index ++
		}else{
			nums[i],nums[i-index] = 0,nums[i]
		}
	}
}

func main()  {
	a := []int{0,1,0,3,12}
	moveZeroes1(a)
	fmt.Println(a)
}
