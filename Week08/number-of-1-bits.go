package main

import "fmt"

//191. 位1的个数
func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		sum++
		num &= (num - 1)
	}
	return sum
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
