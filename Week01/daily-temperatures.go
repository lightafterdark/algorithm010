package main

import "fmt"

//739. 每日温度


/**
暴力法
外循环获取当前下标温度
内循环判断后面是否有高于当前下标温度的，无，则为0，有，则为下标差
 */
func dailyTemperatures1(T []int) []int {
	res := make([]int,0)
	for i:=0;i<len(T);i++  {
		flag := false
		for j:=i+1;j<len(T);j++ {
			if T[j] > T[i] {
				res = append(res,j-i)
				flag = true
				break
			}
		}
		if !flag {
			res = append(res,0)
		}
	}
	return res
}



/**
单调栈
可以维护一个存储下标的单调栈，从栈底到栈顶的下标对应的温度列表中的温度依次递减。
如果一个下标在单调栈里，则表示尚未找到下一次温度更高的下标。
正向遍历温度列表。对于温度列表中的每个元素 T[i]，如果栈为空，则直接将 i 进栈，
如果栈不为空，则比较栈顶元素 prevIndex 对应的温度 T[prevIndex] 和当前温度 T[i]，
如果 T[i] > T[prevIndex]，则将 prevIndex 移除，并将 prevIndex 对应的等待天数赋为 i - prevIndex，
重复上述操作直到栈为空或者栈顶元素对应的温度小于等于当前温度，然后将 i 进栈。
 */
func dailyTemperatures(T []int) []int {
	res := make([]int,len(T))
	stack := []int{}
	for i:=0;i<len(T);i++  {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prevIndex] = i - prevIndex
		}
		stack = append(stack,i)
	}
	return res
}

func main()  {
	test := []int{73, 74, 75, 71, 69, 72, 76, 73}
	test1 := []int{55,38,53,81,61,93,97,32,43,78}
	fmt.Println(dailyTemperatures(test))
	fmt.Println(dailyTemperatures(test1))
}