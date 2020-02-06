package ternary

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTernary(t *testing.T) {
	Convey("exec cond", t, func() {
		op1 := Ternary(func() bool { return true }, "10", "2").AsString()
		So(op1, ShouldEqual, "10")

		op2 := Ternary(func() {}, "10", "2").AsString()
		So(op2, ShouldEqual, "2")

		op3 := Ternary(func(param string) bool { return true }, "10", "2").AsString()
		So(op3, ShouldNotEqual, "10")

		op4 := Ternary(func(param string) (bool, string) { return true, "3" }, "10", "2").AsString()
		So(op4, ShouldNotEqual, "10")

		op5 := Ternary(func() string { return "1" }, "10", "2").AsString()
		So(op5, ShouldNotEqual, "10")

		op6 := Ternary("1", "10", "2").AsString()
		So(op6, ShouldNotEqual, "10")

		op7 := Ternary(func() bool { return false }, "10", "2").AsString()
		So(op7, ShouldEqual, "2")
	})
}
