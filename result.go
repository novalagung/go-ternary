package ternary

import (
	"fmt"
	"reflect"
	"strconv"
)

// Result is the type returned from Ternary() function
type Result struct {
	value interface{}
}

// AsBool cast result as bool
func (r Result) AsBool() bool {
	switch r.value.(type) {
	case bool:
		return r.value.(bool)
	default:
		return !reflect.ValueOf(r.value).IsZero()
	}
}

// AsInt cast result as int
func (r Result) AsInt() int {
	switch r.value.(type) {
	case int:
		return r.value.(int)
	default:
		n, _ := strconv.ParseInt(r.AsString(), 10, 32)
		return int(n)
	}
}

// AsInt8 cast result as int8
func (r Result) AsInt8() int8 {
	switch r.value.(type) {
	case int8:
		return r.value.(int8)
	default:
		n, _ := strconv.ParseInt(r.AsString(), 10, 8)
		return int8(n)
	}
}

// AsInt16 cast result as int16
func (r Result) AsInt16() int16 {
	switch r.value.(type) {
	case int16:
		return r.value.(int16)
	default:
		n, _ := strconv.ParseInt(r.AsString(), 10, 16)
		return int16(n)
	}
}

// AsInt32 cast result as int32
func (r Result) AsInt32() int32 {
	switch r.value.(type) {
	case int32:
		return r.value.(int32)
	default:
		n, _ := strconv.ParseInt(r.AsString(), 10, 32)
		return int32(n)
	}
}

// AsInt64 cast result as int64
func (r Result) AsInt64() int64 {
	switch r.value.(type) {
	case int64:
		return r.value.(int64)
	default:
		n, _ := strconv.ParseInt(r.AsString(), 10, 64)
		return int64(n)
	}
}

// AsUint8 cast result as uint8
func (r Result) AsUint8() uint8 {
	switch r.value.(type) {
	case uint8:
		return r.value.(uint8)
	default:
		n, _ := strconv.ParseUint(r.AsString(), 10, 8)
		return uint8(n)
	}
}

// AsUint16 cast result as uint16
func (r Result) AsUint16() uint16 {
	switch r.value.(type) {
	case uint16:
		return r.value.(uint16)
	default:
		n, _ := strconv.ParseUint(r.AsString(), 10, 16)
		return uint16(n)
	}
}

// AsUint32 cast result as uint32
func (r Result) AsUint32() uint32 {
	switch r.value.(type) {
	case uint32:
		return r.value.(uint32)
	default:
		n, _ := strconv.ParseUint(r.AsString(), 10, 32)
		return uint32(n)
	}
}

// AsUint64 cast result as uint64
func (r Result) AsUint64() uint64 {
	switch r.value.(type) {
	case uint64:
		return r.value.(uint64)
	default:
		n, _ := strconv.ParseUint(r.AsString(), 10, 64)
		return uint64(n)
	}
}

// AsUintptr cast result as uintptr
func (r Result) AsUintptr() uintptr {
	switch r.value.(type) {
	case uintptr:
		return r.value.(uintptr)
	default:
		n, _ := strconv.ParseUint(r.AsString(), 10, 64)
		return uintptr(n)
	}
}

// AsFloat32 cast result as float32
func (r Result) AsFloat32() float32 {
	switch r.value.(type) {
	case float32:
		return r.value.(float32)
	default:
		n, _ := strconv.ParseFloat(r.AsString(), 32)
		return float32(n)
	}
}

// AsFloat64 cast result as float64
func (r Result) AsFloat64() float64 {
	switch r.value.(type) {
	case float64:
		return r.value.(float64)
	default:
		n, _ := strconv.ParseFloat(r.AsString(), 32)
		return float64(n)
	}
}

// AsComplex64 cast result as complex64
func (r Result) AsComplex64() complex64 {
	switch r.value.(type) {
	case complex64:
		return r.value.(complex64)
	default:
		return complex64(complex(r.AsFloat64(), 0))
	}
}

// AsComplex128 cast result as complex128
func (r Result) AsComplex128() complex128 {
	switch r.value.(type) {
	case complex128:
		return r.value.(complex128)
	default:
		return complex(r.AsFloat64(), 0)
	}
}

// AsString cast result as string
func (r Result) AsString() string {
	switch r.value.(type) {
	case string:
		return r.value.(string)
	default:
		return fmt.Sprintf("%v", r.value)
	}
}

// AsInterface cast result as string
func (r Result) AsInterface() interface{} {
	return r.value
}

// StoreTo cast result as string
func (r Result) StoreTo(val interface{}) {
	defer (func() { recover() })()

	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Ptr {

		el := rv.Elem()
		if el.CanSet() {
			el.Set(reflect.ValueOf(r.value))
		}
	}
}

// ExecIfResultIsFunc is used to call the truthy and falsey argument that are actually a func
func (r Result) ExecIfResultIsFunc() Result {
	if isValueValidExecutableFunc(r.value) {
		res := reflect.ValueOf(r.value).Call(make([]reflect.Value, 0))[0]
		return Result{value: res.Interface()}
	}

	return r
}

func isValueValidExecutableFunc(any interface{}) bool {
	rt := reflect.TypeOf(any)
	if rt.Kind() == reflect.Func {
		if rt.NumIn() > 0 {
			return false
		}
		if rt.NumOut() != 1 {
			return false
		}

		return true
	}

	return false
}
