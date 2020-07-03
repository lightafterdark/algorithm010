package main

import "fmt"

//860. 柠檬水找零
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for _,bill := range bills {
		if bill == 5 {
			five += 1
		}else if bill == 10 {
			if five < 1 {
				return false
			}
			five -= 1
			ten += 1
		}else if bill == 20 {
			if five >= 1 && ten >= 1 {
				five -= 1
				ten -= 1
			}else if five >= 3 {
				five -= 3
			}else {
				return false
			}
		}else{
			return false
		}
	}
	return true
}

func main()  {
	fmt.Println(lemonadeChange([]int{5,5,10,20}))
	fmt.Println(lemonadeChange([]int{5,5,10}))
	fmt.Println(lemonadeChange([]int{10,20}))
}