package main

import "fmt"

//190. 颠倒二进制位
func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := uint32(31); num > 0; i-- {
		res = res + (num&uint32(1))<<i
		num >>= 1
	}
	return res
}
func main() {
	fmt.Println(reverseBits(00000010100101000001111010011100))

}
