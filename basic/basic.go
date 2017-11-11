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
	for n > 1 {
		fact *= n
		n--
	}
	return fact
}

//IsPalindrome checks if a string can be read in both directions, omitting digits and non-letter characters
func IsPalindrome(str string) bool {
	isPalindrome := true
	// 	front := 0
	// 	back := len(str) - 1
	// 	re := regexp.MustCompile(`[\W_]`)
	// 	for front > back {
	// for front, substr := range str {
	// 	re.MatchString(substr)
	// }

	// 		for re.MatchString(string(str[front])) {
	// 			front++
	// 			continue
	// 		}

	// 		for re.MatchString(string(str[back])) {
	// 			back--
	// 			continue
	// 		}

	// 		if str[front] != str[back] {
	// 			isPalindrome = false
	// 			break
	// 		}
	// 		front++
	// 		back--
	// 	}
	return isPalindrome
}

func FindLongestWord(str string) int {
	return 0
}

func TitleCase(str string) string {
	return ""
}

func LargestOfFour([4][4]int) [4]int {
	return [4]int{0, 0, 0, 0}
}

func ConfirmEnding(str string, ending string) bool {
	return false
}

func RepeatStringNumTimes(str string, n int) string {
	return ""
}

func TruncateString(str string, n int) string {
	return ""
}

func ChunkArrayInGroups(arr []int, size int) [][]int {
	return [][]int{}
}
