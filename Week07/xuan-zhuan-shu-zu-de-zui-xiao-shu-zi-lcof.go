package main

import "fmt"

//func minArray(numbers []int) int {
//	if len(numbers) < 1 {
//		return 0
//	}
//	res := numbers[0]
//	for _, num := range numbers {
//		if num < res {
//			res = num
//		}
//	}
//	return res
//}

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		middle := (left + right) / 2
		if numbers[middle] < numbers[right] {
			right = middle
		} else if numbers[middle] > numbers[right] {
			left = middle + 1
		} else {
			right -= 1
		}
	}
	return numbers[left]
}
func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}
