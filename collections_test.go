package gogenutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	lst := []string{"haha", "hihi", "hoho"}
	s := "hihi"
	assert.True(t, InSlice(s, lst), fmt.Sprintf("%s in %v", s, lst))

	lst = []string{"hahaha", "hihi", "hoho"}
	s = "haha"
	assert.False(t, InSlice(s, lst), fmt.Sprintf("%s in %v", s, lst))

	ilst := []int{12, 45, 76, 32}
	i := 34
	assert.False(t, InSlice(i, ilst), fmt.Sprintf("%d in %v", i, ilst))

	i = 45
	assert.True(t, InSlice(i, ilst), fmt.Sprintf("%d in %v", i, ilst))
}

func TestSlicesEqual(t *testing.T) {
	a := []string{"haha", "hihi", "hoho"}
	b := []string{"haha", "hihi", "hoho"}
	assert.True(t, SlicesEqual(a, b), fmt.Sprintf("%v equals %v (element wise)", a, b))
	b = []string{"haha", "hiha", "hoho"}
	assert.False(t, SlicesEqual(a, b), fmt.Sprintf("%v ! equals %v (element wise)", a, b))
}

func TestFilterCommonRootDirs(t *testing.T) {
	dirs := func() []string { // haha, we need to do this to make sure dirs is a constant!!!
		return []string{"/var", "/home/john", "/usr/lib/", "/var/lib", "/home/john/Documents", "/usr/lib"}
	}
	expected := []string{"/var", "/home/john", "/usr/lib"}
	filtered := FilterCommonRootDirs(dirs())
	cmpres := SlicesEqual(expected, filtered)
	assert.True(t, cmpres, fmt.Sprintf("Filtered %v => %v ?= %v", dirs(), filtered, expected))
}
