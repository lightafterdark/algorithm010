package main

import "fmt"

//20.有效得括号
/**
左括号则入栈
否则
 */
func isValid(s string) bool {
	dic := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)

	for i:=0;i<len(s) ;i++  {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack,s[i])
		}else if len(stack) > 0 && stack[len(stack) - 1] == dic[s[i]] {
			stack = stack [:len(stack)-1]
		}else{
			return false
		}
	}

	return len(stack)  == 0
}

func main()  {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid( "{[]}"))
	fmt.Println(isValid("(]"))
}