package main
//49. 字母异位词分组
import (
	"fmt"
	"sort"
	"strings"
)

type SortedString struct {
	Oring string
	Sorted string
}
func groupAnagrams1(strs []string) [][]string {
	res := make([][]string,0)
	items := make([]*SortedString,len(strs))
	for i,str := range strs {
		sortedValue := strings.Split(str,"")
		sort.Slice(sortedValue, func(a, b int) bool {
			if sortedValue[a]>sortedValue[b]{
				return true
			}else{
				return false
			}
		})
		items[i] = &SortedString{
			Oring:str,
			Sorted:strings.Join(sortedValue,""),
		}
	}
	sort.Slice(items, func(a, b int) bool {
		if items[a].Sorted > items[b].Sorted {
			return true
		}else{
			return false
		}
	})

	for len(items) > 0 {
		curItem := items[0]
		items = items[1:]
		temp := []string{curItem.Oring}
		for len(items)>0 && curItem.Sorted == items[0].Sorted {
			temp = append(temp,items[0].Oring)
			items = items[1:]
		}
		res = append(res,temp)
	}
	return res
}


type bytes []byte
// 自定义排序规则
func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i,j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string]int)
	for _,str := range strs {
		kByte := bytes(str)
		sort.Sort(kByte)
		k := string(kByte)
		if idx,ok := m[k];!ok {
			m[k] = len(res)
			res = append(res,[]string{str})
		}else{
			res[idx] = append(res[idx],str)
		}
	}
	return res
}

func main()  {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}