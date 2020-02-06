package ternary

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResult(t *testing.T) {
	Convey("cast result as bool", t, func() {
		So(Ternary(true, true, 2).AsBool(), ShouldEqual, true)
		So(Ternary(true, int32(0), 2).AsBool(), ShouldEqual, false)
		So(Ternary(true, float64(0), 2).AsBool(), ShouldEqual, false)
	})

	Convey("cast result to int", t, func() {
		So(Ternary(true, int(10), 2).AsInt(), ShouldEqual, int(10))
		So(Ternary(true, float64(10), 2).AsInt(), ShouldEqual, int(10))

		So(Ternary(true, int8(10), 2).AsInt8(), ShouldEqual, int8(10))
		So(Ternary(true, float64(10), 2).AsInt8(), ShouldEqual, int8(10))

		So(Ternary(true, int16(10), 2).AsInt16(), ShouldEqual, int16(10))
		So(Ternary(true, float64(10), 2).AsInt16(), ShouldEqual, int16(10))

		So(Ternary(true, int32(10), 2).AsInt32(), ShouldEqual, int32(10))
		So(Ternary(true, float64(10), 2).AsInt32(), ShouldEqual, int32(10))

		So(Ternary(true, int64(10), 2).AsInt64(), ShouldEqual, int64(10))
		So(Ternary(true, float64(10), 2).AsInt64(), ShouldEqual, int64(10))
	})

	Convey("cast result to uint", t, func() {
		So(Ternary(true, uint8(10), 2).AsUint8(), ShouldEqual, uint8(10))
		So(Ternary(true, int64(10), 2).AsUint8(), ShouldEqual, uint8(10))

		So(Ternary(true, uint16(10), 2).AsUint16(), ShouldEqual, uint16(10))
		So(Ternary(true, int64(10), 2).AsUint16(), ShouldEqual, uint16(10))

		So(Ternary(true, uint32(10), 2).AsUint32(), ShouldEqual, uint32(10))
		So(Ternary(true, int64(10), 2).AsUint32(), ShouldEqual, uint32(10))

		So(Ternary(true, uint64(10), 2).AsUint64(), ShouldEqual, uint64(10))
		So(Ternary(true, int64(10), 2).AsUint64(), ShouldEqual, uint64(10))

		So(Ternary(true, uintptr(10), 2).AsUintptr(), ShouldEqual, uintptr(10))
		So(Ternary(true, int64(10), 2).AsUintptr(), ShouldEqual, uintptr(10))
	})

	Convey("cast result to float", t, func() {
		So(Ternary(true, float32(10), 2).AsFloat32(), ShouldEqual, float32(10))
		So(Ternary(true, int64(10), 2).AsFloat32(), ShouldEqual, float32(10))

		So(Ternary(true, float64(10), 2).AsFloat64(), ShouldEqual, float64(10))
		So(Ternary(true, int64(10), 2).AsFloat64(), ShouldEqual, float64(10))
	})

	Convey("cast result to complex", t, func() {
		So(Ternary(true, complex64(10), 2).AsComplex64(), ShouldEqual, complex64(10))
		So(Ternary(true, int64(10), 2).AsComplex64(), ShouldEqual, complex64(10))

		So(Ternary(true, complex128(10), 2).AsComplex128(), ShouldEqual, complex64(10))
		So(Ternary(true, int64(10), 2).AsComplex128(), ShouldEqual, complex64(10))
	})

	Convey("cast result to string", t, func() {
		So(Ternary(true, "10", 2).AsString(), ShouldEqual, "10")
		So(Ternary(true, int(10), 2).AsString(), ShouldEqual, "10")
		So(Ternary(true, float32(10), 2).AsString(), ShouldEqual, "10")
	})

	Convey("cast result to interface{}", t, func() {
		So(Ternary(true, "10", 2).AsInterface().(string), ShouldEqual, "10")
		So(Ternary(true, int(10), 2).AsInterface().(int), ShouldEqual, int(10))
	})

	Convey("store result to pointer var", t, func() {
		var holder1 string
		Ternary(true, "10", 2).StoreTo(&holder1)
		So(holder1, ShouldEqual, "10")

		var holder2 string
		Ternary(true, "10", 2).StoreTo(holder2)
		So(holder2, ShouldEqual, "")
	})

	Convey("exec result", t, func() {
		op1 := Ternary(true, func() string { return "10" }, func() string { return "2" }).ExecIfResultIsFunc().AsString()
		So(op1, ShouldEqual, "10")

		op2 := Ternary(false, "10", func() string { return "2" }).ExecIfResultIsFunc().AsString()
		So(op2, ShouldEqual, "2")

		op3 := Ternary(false, "10", func(param string) string { return "2" }).ExecIfResultIsFunc().AsString()
		So(op3, ShouldNotEqual, "10")

		op4 := Ternary(false, "10", func() (string, string) { return "2", "2" }).ExecIfResultIsFunc().AsString()
		So(op4, ShouldNotEqual, "10")

		op5 := Ternary(false, "10", "2").ExecIfResultIsFunc().AsString()
		So(op5, ShouldNotEqual, "10")
	})
}
