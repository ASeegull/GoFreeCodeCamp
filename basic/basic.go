package basic

import (
	"errors"
)

// ****** Reverse String *******
//
// You need to reverse the provided string.
// You may need to turn the string into an
// array before you can reverse it.
// And your result must be a string.
//

func ReverseString(str string) string {
	r := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		r[i] = str[len(str)-i-1]
	}
	return string(r)
}

// ******** Factorialize Number ********
// Take an int and return factorial of it
// Just for now there is no need to check for errors
//

func Factorialize(n int) int {
	fact := 1
	for n > 1 {
		fact *= n
		n--
	}
	return fact
}

// ******** Factorialize With Errors ********
//
// Now let's try to do the same function, but check for errors
// Take a look at errors package and make brand New error
func FactorializeWithError(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Can't factorialize negative number")
	}
	fact := 1
	for n > 1 {
		fact *= n
		n--
	}
	return fact, nil
}

// ******** Check for Palindromes ********
//
// Return true if the given string is a palindrome. Otherwise, return false.
// A palindrome is a word or sentence that's spelled the same way both forward and backward,
// ignoring punctuation, case, and spacing.
//
// Note!
//
// You'll need to remove all non-alphanumeric characters (punctuation, spaces and symbols)
// and turn everything lower case in order to check for palindromes.
//

func IsPalindrome(str string) bool {
	// isPalindrome := true
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
	return false
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
