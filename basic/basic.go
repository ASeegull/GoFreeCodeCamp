package basic

import (
	"fmt"
)

func main() {
	fmt.Println(ReverseString("hello"))
}

//ReverseString takes a string and returns it's reversed variant
func ReverseString(str string) string {
	r := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		r[i] = str[len(str)-i-1]
	}
	return string(r)
}

//Factorialize takes an int and returns factorial of it
func Factorialize(n int) int {
	fact := 1
	for n > 0 {
		fact *= n
		n--
	}
	return fact
}

func IsPalindrome(str string) bool {

	return false
}
