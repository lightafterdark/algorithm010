package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	n1 := m - 1
	n2 := n - 1
	l := m + n - 1
	for n2 >= 0 {
		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[l] = nums1[n1]
			n1--
		} else {
			nums1[l] = nums2[n2]
			n2--
		}
		l--
	}
	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 4, 6}, 3)

}
