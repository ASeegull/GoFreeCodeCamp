package factorialize

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFactorialize(t *testing.T) {
	Convey("It should return correct value of factorial", t, func() {

		Convey("Factorialize `5` Should return => `120`", func() {
			test1 := Factorialize(5)
			So(test1, ShouldEqual, 120)
		})

		Convey("Factorialize `10` Should return => `3628800`", func() {
			test2 := Factorialize(10)
			So(test2, ShouldEqual, 3628800)
		})

		Convey("Factorialize `20` Should return => `2432902008176640000`", func() {
			test3 := Factorialize(20)
			So(test3, ShouldEqual, 2432902008176640000)
		})

		Convey("Factorialize `0` Should return =>  `1`", func() {
			test4 := Factorialize(0)
			So(test4, ShouldEqual, 1)
		})
	})
}

func TestFactorializeWithErrors(t *testing.T) {
	Convey("It should return correct value of factorial and an error", t, func() {

		Convey("Factorialize `5` Should return => `120`", func() {
			test1, err := FactorializeWithError(5)
			So(test1, ShouldEqual, 120)
			So(err, ShouldBeNil)
		})

		Convey("Factorialize `10` Should return => `3628800`", func() {
			test2, err := FactorializeWithError(10)
			So(test2, ShouldEqual, 3628800)
			So(err, ShouldBeNil)
		})

		Convey("Factorialize `20` Should return => `2432902008176640000`", func() {
			test3, err := FactorializeWithError(20)
			So(test3, ShouldEqual, 2432902008176640000)
			So(err, ShouldBeNil)
		})

		Convey("Factorialize `0` Should return =>  `1`", func() {
			test4, err := FactorializeWithError(0)
			So(test4, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Factorialize `-6` Should return =>  `0` and throw error", func() {
			test5, err := FactorializeWithError(-6)
			So(test5, ShouldEqual, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
