package main

import "fmt"

//14. 最长公共前缀




func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i:=0;i<len(strs[0]);i++  {
		//外循环为第一个字符串的元素
		for j:=1;j<len(strs);j++{
			//内循环为除了第一个字符串外所有的元素
			if i == len(strs[j]) || strs[j][i] != strs[0][i]{
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main()  {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}