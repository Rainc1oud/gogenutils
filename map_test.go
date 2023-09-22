package gogenutils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type tMyElem struct {
	SomeString      string
	SomeInt         int
	SomeOtherString string
}

func TestFilters(t *testing.T) {
	m := NewMap[string, tMyElem]()
	m["item1"] = tMyElem{SomeString: "item1-string", SomeInt: 3}
	m["item2"] = tMyElem{SomeString: "item2-string", SomeInt: 33}
	m["item3"] = tMyElem{SomeString: "item3-string", SomeInt: 6}
	m["item4"] = tMyElem{SomeString: "item4-string", SomeInt: 9}

	fm1 := m.FilterMapByKey(func(s string) bool { return strings.HasPrefix(s, "item2") })
	em1 := NewMap[string, tMyElem]()
	em1["item2"] = m["item2"]
	require.EqualValues(t, em1, fm1)

	fm2 := m.FilterMapByVal(func(e tMyElem) bool { return e.SomeInt >= 9 })
	em2 := Map[string, tMyElem]{"item2": m["item2"], "item4": m["item4"]}
	require.EqualValues(t, em2, fm2)

	// and now for some composition
	fm3 := m.
		FilterMapByVal(func(e tMyElem) bool { return e.SomeInt > 5 }).
		FilterMapByKey(func(s string) bool { return s == "item2" || s == "item3" })
	em3 := Map[string, tMyElem]{"item2": m["item2"], "item3": m["item3"]}
	require.EqualValues(t, em3, fm3)

	// TODO: is there realy no direct way???
	is := TypedSlice[int](m.AttrSlice(func(e tMyElem) any { return e.SomeInt }))
	es := []int{3, 33, 6, 9}
	require.EqualValues(t, es, is)
}
