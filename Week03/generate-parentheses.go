package main
////22. 括号生成
import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string,0)
	generate(0,0,n,"",&res)
	return res
}

func generate(left,right,n int,s string,res *[]string){
	if left == n && right == n {
		*res = append(*res,s)
	}

	if left < n {
		generate(left+1,right,n,s+"(",res)
	}

	if left > right {
		generate(left,right+1,n,s+")",res)
	}
}

func main()  {
	fmt.Println(generateParenthesis(3))
}