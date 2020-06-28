package main

import "fmt"

//17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	res := make([]string,0)
	if digits == "" {
		return res
	}

	dict := map[string][]string{
		"2":{"a","b","c"},
		"3":{"d","e","f"},
		"4":{"g","h","i"},
		"5":{"j","k","l"},
		"7":{"m","n","o"},
		"8":{"p","q","r","s"},
		"9":{"a","b","c"},
	}

	res = append(res,"")

	for i:=0 ; i < len(digits) ; i++  {
		retLen := len(res)
		digit := digits[i:i+1]
		letterList := dict[digit]
		for j:=0;j<retLen;j++ {
			for _,letter := range letterList {
				res = append(res,res[0]+letter)
			}
			res = res[1:]
		}
		
	}
	return res
}

func main()  {
	fmt.Println(letterCombinations("23"))

}

