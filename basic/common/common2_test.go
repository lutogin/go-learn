package common

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestFactorialByConvey(t *testing.T) {
	x := 0
	Convey("A", t, func() {
		x++ // best place for se tup env
		So(1, ShouldEqual, 1)
		So(2*2, ShouldEqual, 4)
		So([]int{1, 2, 3}, ShouldContain, 1)
		So(map[string]int{"test": 1}, ShouldContainKey, "test")
		So(func() { panic("a") }, ShouldPanic)
		So(time.Now().Truncate(10*time.Minute), ShouldHappenBefore, time.Now())

		Convey("A-B", func() { // argument T shouldn't be passed
			x++
			Convey("A-B-C", func() { // argument T shouldn't be passed
				So(x, ShouldEqual, 2)
			})
			Convey("A-B-C1", func() { // argument T shouldn't be passed
				So(x, ShouldEqual, 4)
			})
		})

		Reset(func() {
			t.Log("finish.")
		})
	})
}
