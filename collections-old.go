package gogenutils

import (
	"os"
	"reflect"
	"strings"
)

// FilterSlicePfxStr filters out all string items of which another slice member is a prefix, recursively
// by convention take the first element as the "pivot"
func FilterSlicePfxStr(s []string) []string {
	if len(s) < 2 {
		return s
	}
	r := FilterSlicePfxStr(s[1:])    // recursive except first element ("pivot")
	r = append([]string{s[0]}, r...) // "insert" s[0] at pos 0 of r
	for i := 1; i < len(r); i++ {    // first item is reserved as pivot (to compare to)
		if strings.HasPrefix(r[i], r[0]) {
			if i < len(r)-1 {
				r = append(r[0:i], r[i+1:]...) // remove r[i] because it has the prefix
			} else {
				r = r[0:i]
			}
		}
	}
	return r
}

// FilterCommonRootDirs recursively filters out paths for which any higher root path exists in the same slice
func FilterCommonRootDirs(strSlice []string) []string {
	// sanitise trailing slashes
	for i, s := range strSlice {
		strSlice[i] = strings.TrimSuffix(s, string(os.PathSeparator))
	}
	return FilterSlicePfxStr(strSlice)
}

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

// RemoveFromStringSlice removes the itmen and
// returns a new slice
func RemoveFromStringSlice(item string, slice []string) []string {
	s := slice

	p := PosInSlice(item, s)
	if p < 0 { // item not found, return original
		return slice
	}
	if p > -1 && p < len(s)-1 { // item between (incl.) first and before-last pos
		return append(s[:p], s[p+1:]...)
	}
	return s[:p] // remaining possibility, item to remove is at the last position
}

// Equal tells whether a and b contain the same elements in the same order (!)
// A nil argument is equivalent to an empty slice.
func SlicesEqual(a, b interface{}) bool {
	ai := reflect.ValueOf(a)
	bi := reflect.ValueOf(b)
	if ai == bi {
		return true
	}
	if ai.Kind() != reflect.Slice || ai.IsNil() || ai.IsZero() {
		return false
	}

	if bi.Kind() != reflect.Slice || bi.IsNil() || bi.IsZero() {
		return false
	}

	if ai.Len() != bi.Len() {
		return false
	}
	for i := 0; i < ai.Len(); i++ {
		if ai.Index(i).Interface() != bi.Index(i).Interface() {
			return false
		}
	}
	return true
}

func StringSliceUniq(sslice *[]string) *[]string {
	elem := make(map[string]struct{})
	result := []string{}

	for _, str := range *sslice {
		if _, ok := elem[str]; !ok {
			elem[str] = struct{}{}
			result = append(result, str)
		}
	}

	return &result
}
