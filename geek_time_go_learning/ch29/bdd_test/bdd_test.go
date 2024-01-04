package bdd_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBdd(t *testing.T) {
	Convey("Give 2 even numbers", t, func() {
		a := 2
		//a := 3
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c % 2, ShouldEqual, 0)
			})
		})
	})
}