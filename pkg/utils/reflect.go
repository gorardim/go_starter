package utils

import "reflect"

func NewPointerType[T any]() T {
	var t T
	rt := reflect.TypeOf(t)
	if rt.Kind() != reflect.Ptr {
		panic("not a pointer type")
	}
	elem := reflect.New(rt.Elem())
	return elem.Interface().(T)
}
