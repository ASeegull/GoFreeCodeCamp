package IsPalindrome

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPalindrome(t *testing.T) {
	Convey(`It should remove all non-word charachters and then return true, 
		if string is pallindrom`, t, func() {

		Convey("`_eye)` Should return => `true`", func() {
			test1 := IsPalindrome("_eye)")
			So(test1, ShouldBeTrue)
		})

		Convey("`race car` Should return => `true`", func() {
			test2 := IsPalindrome("race car")
			So(test2, ShouldBeTrue)
		})

		Convey("`A man, a plan, a canal. Panama` Should return => `true`", func() {
			test3 := IsPalindrome("A man, a plan, a canal. Panama")
			So(test3, ShouldBeTrue)
		})

		Convey("`nope` Should return =>  `false`", func() {
			test4 := IsPalindrome("nope")
			So(test4, ShouldBeFalse)
		})

		Convey("`not a palindrome` Should return =>  `false`", func() {
			test5 := IsPalindrome("not a palindrome")
			So(test5, ShouldBeFalse)
		})

		Convey("`1 eye for of 1 eye.` Should return =>  `false`", func() {
			test5 := IsPalindrome("1 eye for of 1 eye.")
			So(test5, ShouldBeFalse)
		})

		// If you are reading this, note that I'm using here backticks, so that go doesn't consider \
		// is an escape character
		Convey(`"0_0 (: /-\ :) 0-0" Should return =>  "true"`, func() {
			test6 := IsPalindrome(`0_0 (: /-\ :) 0-0`)
			So(test6, ShouldBeTrue)
		})

		Convey(`"five|\_/|four" Should return =>  "false"`, func() {
			test6 := IsPalindrome(`five|\_/|four`)
			So(test6, ShouldBeFalse)
		})
	})
}
