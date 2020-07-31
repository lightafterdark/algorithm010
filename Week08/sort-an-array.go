package main

import "fmt"

//912. 排序数组

//选择排序排序放入新数组
//func sortArray1(nums []int) []int {
//	res := []int{}
//	for len(nums) > 0 {
//		min := nums[0]
//		index := 0
//		for i, num := range nums {
//			if num < min {
//				min = num
//				index = i
//			}
//		}
//		res = append(res, min)
//		nums = append(nums[:index], nums[index+1:]...)
//	}
//	return res
//}

//选择排序操作原始数组
func sortArray1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				index = j
			}
		}
		if i != index {
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	return nums
}

//冒泡排序
func sortArray2(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

//插入排序
func sortArray3(nums []int) []int {
	//外层是当次排序需要排的num
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

type Heap struct {
	arr  []int
	size int
}

//堆排序
func sortArray4(nums []int) []int {

	return nums
}

func adjustHeap(h Heap, parentNode int) {
	leftNode := parentNode*2 + 1
	rightNode := parentNode*2 + 2
	maxNode := parentNode
	if leftNode < h.size && h.arr[maxNode] < h.arr[leftNode] {
		maxNode = leftNode
	}
	if rightNode < h.size && h.arr[maxNode] < h.arr[rightNode] {
		maxNode = rightNode
	}

	if maxNode != parentNode {
		h.arr[maxNode], h.arr[parentNode] = h.arr[parentNode], h.arr[maxNode]
		adjustHeap(h, maxNode)
	}
}

func createHeap(arr []int) (h Heap) {
	h.arr = arr
	h.size = len(arr)
	for i := h.size / 2; i >= 0; i-- {
		adjustHeap(h, i)
	}
	return
}

//归并排序
func sortArray5(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2
	left := sortArray5(nums[:mid])
	right := sortArray5(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) (res []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}

//快递排序
//func sortArray(nums []int) []int {
//	left := 0
//	right := len(nums)
//	for left < right {
//		loc := partition(nums, left, right)
//		sortArray(nums)
//
//	}
//	return nums
//}
//
//func partition(nums []int, left, right int) int {
//	i := left + 1
//	j := right
//	for i < j {
//		if nums[i] > nums[left] {
//			nums[i], nums[j] = nums[j], nums[i]
//			j--
//		} else {
//			i++
//		}
//	}
//	if nums[i] >= nums[left] {
//		i--
//	}
//	nums[left], nums[i] = nums[i], nums[left]
//	return i
//}

func main() {
	fmt.Println(sortArray5([]int{5, 2, 3, 1}))
}
