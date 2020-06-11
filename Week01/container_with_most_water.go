package main

import "fmt"

//11.盛最多水得容器
func maxArea(height []int) int {
	max := 0
	l := 0
	r := len(height) -1
	for l < r {
		area := 0
		if height[l] < height[r] {
			area = (r-l)*height[l]
			l ++
		}else{
			area = (r-l)*height[r]
			r --
		}
		if area > max {
			max = area
		}
	}

	return max
}

func main()  {
	a := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(a))
}