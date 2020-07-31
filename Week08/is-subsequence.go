package main

import "fmt"

func isSubsequence(s string, t string) bool {
	m := len(s)
	n := len(t)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == m
}

func main() {
	fmt.Println(isSubsequence("abf", "abcdef"))
	fmt.Println(isSubsequence("afe", "abcdef"))

}
