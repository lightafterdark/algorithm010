package main
//84.柱状图中最大的矩形
import "fmt"


/**
暴力法一：枚举宽
TODO
*/
func largestRectangleArea2(heights []int) int {
	if len(heights) < 1 {
		return 0
	}else if len(heights) == 1 {
		return heights[0]*1
	}
	max := 0
	for left:=0;left<=len(heights)-1;left++  {
		for right:=left+1; right<=len(heights)-1;right++  {
			area := 0
			if heights[left]>heights[right]{
				area = heights[right]*(right-left+1)
			}else{
				area = heights[left]*(right-left+1)
			}
			fmt.Println(left,right,heights[left],heights[right],area)
			if area > max {
				max = area
			}
		}
	}
	return max
}

/**
暴力法2：枚举高
 */
func largestRectangleArea1(heights []int) int {
	max := 0
	if len(heights) < 1 {
		return 0
	}else if len(heights) == 1 {
		return heights[0]*1
	}

	for i:=0;i<=len(heights)-1;i++  {
		height := heights[i]
		left := i
		right := i
		for left -1 >= 0 && heights[left-1] >= height {
			left --
		}
		for right + 1 < len(heights) && heights[right+1] >= height  {
			right ++
		}
		area := (right-left+1)*height
		if area > max {
			max = area
		}
	}
	return max
}


/**
单调栈
利用一个栈实现
这个栈用来存储最大高度得下标
遍历数组

 */
func largestRectangleArea(heights []int) int {
	res := 0
	heights = append([]int{-2},heights...)
	heights = append(heights,-1)
	n := len(heights)
	//递增栈 记录最大高度得下标
	s := make([]int,1,n)
	i := 1
	for i < n {
		if heights[s[len(s)-1]] < heights[i] {
			s = append(s,i)
			i++
			continue
		}
		res = max(res,heights[s[len(s)-1]]*(i-s[len(s)-2]-1))
		s = s[:len(s)-1]
	}
	return res
}

func max(a,b int)int{
	if a>b{return a}
	return b
}

func main()  {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println(largestRectangleArea([]int{1,1}))
	fmt.Println(largestRectangleArea([]int{5,1}))
	fmt.Println(largestRectangleArea([]int{1}))
	fmt.Println(largestRectangleArea([]int{0,9}))
}