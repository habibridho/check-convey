package calculation

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("Given two values", t, func() {
		x := 2
		y := 5

		Convey("When it is added", func() {
			result := Add(x, y)

			Convey("Then it should return 7", func() {
				So(result, ShouldEqual, 7)
			})
		})
	})

	Convey("Given two values", t, func() {
		x := 2
		_ = 5

		Convey("When Add is called with not enough argument", func() {
			result := Add(x)

			Convey("Then it should break", func() {
				So(result, ShouldEqual, 2)
			})
		})
	})
}
