package reversestring

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	Convey("It should return the provided string in reversed order", t, func() {

		Convey("`Hello, world!` Should return => `!dlrow ,olleH`", func() {
			test1 := ReverseString("Hello, world!")
			So(test1, ShouldEqual, "!dlrow ,olleH")
		})

		Convey("`Greetings from Earth` Should return => `!dlrow ,olleH`", func() {
			test2 := ReverseString("Greetings from Earth")
			So(test2, ShouldEqual, "htraE morf sgniteerG")
		})

		Convey("`567329-2349` Should return => `9432-923765`", func() {
			test3 := ReverseString("567329-2349")
			So(test3, ShouldEqual, "9432-923765")
		})

		Convey("`GoFreeCodeCamp` Should return =>  `pmaCedoCeerFoG`", func() {
			test4 := ReverseString("GoFreeCodeCamp")
			So(test4, ShouldEqual, "pmaCedoCeerFoG")
		})
	})
}
