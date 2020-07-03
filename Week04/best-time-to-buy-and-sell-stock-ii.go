package main

import "fmt"

//122. 买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	profit := 0
	for i:=1;i<=len(prices)-1;i++  {
		tmp := prices[i]-prices[i-1]
		if tmp > 0 {
			profit += tmp
		}
	}
	return profit
}

func main()  {
	//fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
}