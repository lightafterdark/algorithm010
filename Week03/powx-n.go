package main
//50. Pow(x, n)
import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}else if n == 1 {
		return x
	}else if n < 0 {
		return 1.0/myPow(x,-n)
	}

	subRes := myPow(x,n/2)
	if n % 2 == 1 {
		return subRes*subRes*x
	}else{
		return subRes*subRes
	}
}

func main()  {
	fmt.Println(myPow(2.0,1))
	fmt.Println(myPow(2.0,3))
	fmt.Println(myPow(2.0,0))
	fmt.Println(myPow(2.0,5))
}
