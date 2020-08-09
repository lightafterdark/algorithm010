package main

import "fmt"

//541. 反转字符串 II

func reverseStr(s string, k int) string {
	str :=[]byte(s)
	i := 0
	for i<len(str) {
		l:=i
		r:=i+k-1
		if r>=len(str){
			r =len(str)-1
		}
		for l<r&&l<len(str){
			str[l],str[r]=str[r],str[l]
			l++
			r--
		}
		i=i+2*k
	}
	return string(str)
}

func main()  {
	fmt.Println(reverseStr("abcdefg",2))
}