package main

import "fmt"

func isValid(s string) bool {
	dict := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)
	for i:=0;i<len(s);i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			//坐括号则入栈
			stack = append(stack,s[i])
		}else if len(stack) > 0 && stack[len(stack)-1] == dict[s[i]] {
			//右括号如果栈里面最上面的和它匹配，则出栈一个左括号
			stack = stack[:len(stack)-1]
		}else{
			return false
		}
	}
	return len(stack) == 0
}

func main()  {
	fmt.Println(isValid("((}"))
	fmt.Println(isValid("((}}"))
	fmt.Println(isValid("()"))
}