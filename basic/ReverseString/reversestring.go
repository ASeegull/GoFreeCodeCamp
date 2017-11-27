package reversestring

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
