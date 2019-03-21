package calculation

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSub(t *testing.T) {
	Convey("Given two values", t, func() {
		x := 5
		y := 2

		Convey("When it is subtracted", func() {
			result := Sub(x, y)

			Convey("Then it should return 3", func() {
				So(result, ShouldEqual, 3)
			})
		})
	})
}
