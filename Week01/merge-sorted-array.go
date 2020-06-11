package main

import (
	"fmt"
	"sort"
)

//88. 合并两个有序数组

func merge1(nums1 []int, m int, nums2 []int, n int)  {
	nums1 = append(nums1[:m],nums2...)
	sort.Ints(nums1)
}


/**
因为两个数组仅为有序的
从前往后的话会需要额外的空间
所以可以两个数组从最后一位开始比较，较大值则为合并之后最后一位
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
	n1 := m - 1//num1最后一位
	n2 := n - 1//num2最后一位
	l := m + n - 1 //合并之后总的数组最后一位
	for n2 >= 0 {
		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[l] = nums1[n1]
			n1--
		}else{
			nums1[l] = nums2[n2]
			n2--
		}
		l--
	}
}

func main()  {
	a := []int{1,2,3,0,0,0}
	b := []int{2,5,6}
	merge(a,3,b,3)
	fmt.Println(a)
}