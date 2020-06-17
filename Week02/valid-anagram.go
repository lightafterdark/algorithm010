package main

import (
	"fmt"
	"reflect"
)

//242. 有效的字母异位词

/**
1.遍历两个字符串，存储到map里面
对比两个map
 */
func isAnagram1(s string, t string) bool {
	sMap := make(map[string]int)
	tMap := make(map[string]int)
	for _,ss := range s {
		if n,ok := sMap[string(ss)];ok{
			sMap[string(ss)] = n + 1
		}else{
			sMap[string(ss)] = 1
		}

	}
	for _,tt := range t {
		if n,ok := tMap[string(tt)];ok{
			tMap[string(tt)] = n + 1
		}else{
			tMap[string(tt)] = 1
		}
	}
	if len(sMap)>len(tMap) {
		for k,v := range sMap {
			if n,ok := tMap[k];ok{
				if n!=v{
					return false
				}
			}else{
				return false
			}
		}
	}else{
		for k,v := range tMap {
			if n,ok := sMap[k];ok{
				if n!=v{
					return false
				}
			}else{
				return false
			}
		}
	}
	return true
}


/**
用一个map
遍历第一个字符串存到map，遍历第二个从中删除对应元素，最后判断是否为空map
 */
func isAnagram2(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	res := make(map[string]int)
	for _,c := range s {
		if n,ok := res[string(c)];ok{
			res[string(c)] = n + 1
		}else{
			res[string(c)] = 1
		}
	}

	for _,c := range t {
		if n,ok := res[string(c)];!ok{
			return false
		}else{
			if n>1{
				res[string(c)] = n-1
			}else{
				delete(res,string(c))
			}
		}
	}
	if len(res) != 0 {
		return false
	}
	return true
}

/**
用两个26位的数组，判断数组是否相等
 */
func isAnagram4(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	wordstable_s := [26]int{}
	wordstable_t := [26]int{}
	for _,c := range s {
		index := c-'a'
		wordstable_s[index]++
	}
	for _,c := range t {
		index := c-'a'
		wordstable_t[index]++
	}
	return wordstable_s == wordstable_t
}


func isAnagram(s string, t string) bool {
	sMap := make(map[string]int)
	tMap := make(map[string]int)
	for _,ss := range s {
		sMap[string(ss)]++
	}
	for _,tt := range t {
		tMap[string(tt)]++
	}
	return reflect.DeepEqual(sMap,tMap)
}

func main()  {
	fmt.Println(isAnagram("a","ab"))
	fmt.Println(isAnagram("acd","abd"))
	fmt.Println(isAnagram("adb","abd"))
}