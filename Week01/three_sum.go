package main
//15.三数之和
import (
	"fmt"
	"sort"
)



//TODO 去重
func threeSum1(nums []int) [][]int {
	res := make([][]int,0)
	for i := 0;i < len(nums) - 2;i++  {
		for j := i + 1;j < len(nums) - 1 ;j++  {
			for k := j + 1; k < len(nums) ; k ++  {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := make([]int,0)
					temp = append(temp,nums[i],nums[j],nums[k])
					res = append(res,temp)
				}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	res := make([][]int,0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for k := 0;k < len(nums) - 2;k++  {
		if nums[k]>0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i := k + 1
		j := len(nums) - 1
		for i < j {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				i++
				continue
			}else if sum > 0 {
				j--
				continue
			}else{
				res = append(res,[]int{nums[i],nums[j],nums[k]})
				for {
					if i < j && nums[j] == nums[j-1] {
						j--
					}else{
							break
					}
				}
				i ++
				j --
			}
		}

	}
	return res
}


func main()  {
	a := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(a))
	fmt.Println(threeSum1(a))

}
