package print

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPrint(t *testing.T) {
	Convey("Given a string", t, func() {
		s := "some string"

		Convey("When print is called", func() {
			result := Prt(s)

			Convey("Then it should return some string", func() {
				So(result, ShouldEqual, "some string")
			})
		})
	})
}
