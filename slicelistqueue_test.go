package gogenutils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSLQ(t *testing.T) {
	ssq := NewSimpleSLQ()
	it := ssq.Append("itemlast")
	ih := ssq.Push("itemfirst")
	assert.EqualValues(t, ssq.Length(), 2)
	for ssq.Next() {
		log.Printf("%v\n", ssq.CurrentItem())
	}
	assert.EqualValues(t, 1, ssq.IndexOf(it))
	assert.EqualValues(t, 0, ssq.IndexOf(ih))
	im := ssq.InsertAt("itemmiddle", 1)
	assert.EqualValues(t, 1, ssq.IndexOf(im))
	assert.EqualValues(t, 3, ssq.Length())
}
