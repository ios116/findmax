package findmax

import (
	"reflect"
)

// FindMax returns a max the elements(interface) in the provided slice.
//
// FindMax panics if the provided interface is not a slice.
func FindMax(slice interface{}, less func(i, j int) bool) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("value isn't slice")
	}
	var max int
	for i := 0; i < s.Len(); i++ {
		if less(max, i) {
			max = i
		}

	}
	return s.Index(max).Interface()
}
