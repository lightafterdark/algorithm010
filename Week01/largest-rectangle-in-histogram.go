package main
//84.柱状图中最大的矩形
import "fmt"


//func largestRectangleArea(heights []int) int {
//	max := 0
//	if len(heights) < 1 {
//		return 0
//	}else if len(heights) == 1 {
//		return heights[0]*1
//	}
//
//	for i:=0;i<=len(heights)-2 ;i++  {
//		for j:=i;j<=len(heights)-1;j++  {
//			area := 0
//			if heights[i]>heights[j]{
//				area = heights[j]*(j-i+1)
//			}else{
//				area = heights[i]*(j-i+1)
//			}
//			if area > max {
//				max = area
//			}
//		}
//	}
//	return max
//}

func largestRectangleArea(heights []int) int {
	max := 0
	for left:=0;left<len(heights)-1 ;left++  {
		for right:=left; right<len(heights)-1;right++  {
			area := 0
			if heights[left]>heights[right]{
				area = heights[right]*(right-left+1)
			}else{
				area = heights[left]*(right-left+1)
			}
			if area > max {
				max = area
			}
		}
	}
	return max
}

func main()  {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println(largestRectangleArea([]int{1,1}))
	fmt.Println(largestRectangleArea([]int{5,1}))
	fmt.Println(largestRectangleArea([]int{}))
}