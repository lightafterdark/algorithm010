package main
//9.回文数
import "fmt"
/**
把这个数字反转，不过反转一半就可以了
 */
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0){
		return false
	}
	reversedNumber := 0
	for x > reversedNumber {
		reversedNumber = reversedNumber * 10 + x % 10
		x /= 10
	}
	return x == reversedNumber || x == reversedNumber / 10
}

func main()  {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(120))
	fmt.Println(isPalindrome(12121))
}