package main

import "fmt"

//哈希
func twoSum1(numbers []int, target int) []int {
	res := make([]int, 0)
	if len(numbers) < 1 {
		return res
	}
	dict := make(map[int]int)
	for k, v := range numbers {
		if v > target && target > 0 {
			break
		}
		if index, ok := dict[target-v]; ok {
			res = append(res, index+1)
			res = append(res, k+1)
			return res
		}
		dict[v] = k
	}
	return res
}

//双指针
func twoSum(numbers []int, target int) []int {
	res := make([]int, 0)
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] == target {
			res = append(res, left+1)
			res = append(res, right+1)
			break
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 9))
}
