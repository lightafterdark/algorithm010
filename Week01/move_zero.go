package main
//283移动零

func moveZeroes(nums []int)  {
	index := int(0)
	for i,val := range nums {
		if val != 0 {
			nums[i],nums[index] = nums[index],nums[i]
		}
	}
}

//func moveZeroes(nums []int)  {
//	index := int(0)
//	for i,_ := range nums {
//		if nums[i] == 0 {
//			index ++
//		}else{
//			nums[i],nums[i-index] = 0,nums[i]
//		}
//	}
//}
