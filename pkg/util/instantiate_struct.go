package util

import (
	"reflect"
)

func InstantiateStruct[T any](i T) T {
	t := reflect.TypeOf(&i).Elem()
	v := reflect.ValueOf(&i).Elem()
	v = reflect.New(v.Type().Elem())
	if v.Type().ConvertibleTo(t) {
		return v.Convert(t).Interface().(T)
	}
	return i
}
