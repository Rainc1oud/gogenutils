package gogenutils

import "reflect"

// InSlice returns true if item is in list and false otherwise
func InSlice(item interface{}, slice interface{}) bool {
	s := reflect.ValueOf(slice)
	ival := reflect.ValueOf(item).Interface()
	if s.Kind() != reflect.Slice || s.IsNil() || s.IsZero() {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == ival {
			return true
		}
	}
	return false
}

// PosInSlice returns the position of item in slice or -1 if not found
func PosInSlice(item interface{}, slice interface{}) int {
	s := reflect.ValueOf(slice)
	ival := reflect.ValueOf(item).Interface()
	if s.Kind() != reflect.Slice || s.IsNil() || s.IsZero() {
		return -1
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == ival {
			return i
		}
	}
	return -1
}
