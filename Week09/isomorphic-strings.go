package main

import "fmt"

//205. 同构字符串
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	smap := make(map[string]string)
	tmap := make(map[string]string)
	for i:=0;i<len(s);i++ {
		if val,ok := smap[s[i:i+1]];ok{
			if val != t[i:i+1] {
				return false
			}
		} else if val,ok := tmap[t[i:i+1]];ok{
			if val != s[i:i+1]{
				return false
			}
		}else{
			smap[s[i:i+1]] = t[i:i+1]
			tmap[t[i:i+1]] = s[i:i+1]
		}
	}

	return true

}

func main()  {
	fmt.Println(isIsomorphic("aba","baa"))
}