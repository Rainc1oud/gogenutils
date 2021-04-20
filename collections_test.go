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
