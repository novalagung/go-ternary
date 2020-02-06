package ternary

import (
	"reflect"
)

// Ternary used for performing inline if-else condition
func Ternary(cond, truthy, falsey interface{}) Result {
	if v, ok := cond.(bool); ok {
		if v {
			return Result{value: truthy}
		}
		return Result{value: falsey}
	}

	if isCondValidExecutableFunc(cond) {
		res := reflect.ValueOf(cond).Call(make([]reflect.Value, 0))[0]
		if res.Kind() == reflect.Bool {
			if res.Bool() {
				return Result{value: truthy}
			}

			return Result{value: falsey}
		}
	}

	return Result{value: falsey}
}

func isCondValidExecutableFunc(any interface{}) bool {
	rt := reflect.TypeOf(any)
	if rt.Kind() == reflect.Func {
		if rt.NumIn() > 0 {
			return false
		}
		if rt.NumOut() != 1 {
			return false
		}
		if rt.Out(0).Kind() != reflect.Bool {
			return false
		}

		return true
	}

	return false
}
