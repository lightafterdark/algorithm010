package main

import (
	"fmt"
	"sort"
)

//455. 分发饼干

/**
贪心问题，优先满足胃口小的小朋友
 */
func findContentChildren(g []int, s []int) int {
	if g == nil  || s == nil {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	gi := 0
	si := 0
	for gi < len(g) && si < len(s) {
		if g[gi] <= s[si]{
			gi ++
			si ++
			res += 1
		}else{
			si += 1
		}
	}
	return res
}

func main()  {
	fmt.Println(findContentChildren([]int{1,2,3},[]int{1,1}))
}