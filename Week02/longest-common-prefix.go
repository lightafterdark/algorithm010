package main
import "fmt"
//14. 最长公共前缀
/**
横向扫描
依次遍历字符串数组中的每个字符串，对于每个遍历到的字符串，更新最长公共前缀，
当遍历完所有的字符串以后，即可得到字符串数组中的最长公共前缀。
 */
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for index,_ := range strs {
		prefix = lcp(prefix,strs[index])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix

}

func lcp(str1,str2 string)  string{
	length := min(len(str1),len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index ++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


/**
纵向扫描
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i:=0;i<len(strs[0]);i++  {
		for j:=1;j<len(strs) ;j++  {
			if i == len(strs[j]) || strs[j][i] != strs[0][i]{
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main()  {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}