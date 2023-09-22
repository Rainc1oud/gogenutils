package gogenutils

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	lst := []string{"haha", "hihi", "hoho"}
	s := "hihi"
	assert.True(t, InSlice(s, lst), fmt.Sprintf("assertion failed: %s in %v", s, lst))

	lst = []string{"hahaha", "hihi", "hoho"}
	s = "haha"
	assert.False(t, InSlice(s, lst), fmt.Sprintf("assertion failed: %s in %v", s, lst))

	ilst := []int{12, 45, 76, 32}
	i := 34
	assert.False(t, InSlice(i, ilst), fmt.Sprintf("assertion failed: %d in %v", i, ilst))

	i = 45
	assert.True(t, InSlice(i, ilst), fmt.Sprintf("assertion failed: %d in %v", i, ilst))
}

func TestSlicesEqual(t *testing.T) {
	a := []string{"haha", "hihi", "hoho"}
	b := []string{"haha", "hihi", "hoho"}
	assert.True(t, SlicesEqual(a, b), fmt.Sprintf("assertion failed: %v == %v (element wise)", a, b))
	b = []string{"haha", "hiha", "hoho"}
	assert.False(t, SlicesEqual(a, b), fmt.Sprintf("assertion failed: %v != %v (element wise)", a, b))
}

func TestFilterCommonRootDirs(t *testing.T) {
	dirs := func() []string { // haha, we need to do this to make sure dirs is a constant!!!
		return []string{"/var", "/home/john", "/usr/lib/", "/var/lib", "/home/john/Documents", "/usr/lib", "/usr/lib/arm-linux", "/some/other/dir/deeper1", "/some/other/dir/deeper2", "/some/other/dir/deeper2/moredeeper4"}
	}
	expected := []string{"/var", "/home/john", "/usr/lib", "/some/other/dir/deeper1", "/some/other/dir/deeper2"}
	filtered := FilterCommonRootDirs(dirs())
	cmpres := SlicesEqual(expected, filtered)
	assert.True(t, cmpres, fmt.Sprintf("assertion failed: filtered %v => %v == %v", dirs(), filtered, expected))
}

func TestPosInSlice(t *testing.T) {
	type args struct {
		item  interface{}
		slice interface{}
	}

	randStrSlice := make([]string, 10)
	for i := 0; i < 10; i += 1 {
		randStrSlice[i] = fmt.Sprintf("item%d", rand.Intn(100))
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "alphanumasc",
			args: args{
				item:  "item3",
				slice: []string{"item0", "item2", "item3", "item4"},
			},
			want: 2,
		},
		{
			name: "alphanumdesc",
			args: args{
				item:  "item3",
				slice: []string{"item4", "item3", "item2", "item1"},
			},
			want: 1,
		},
		{
			name: "randstr",
			args: args{
				item:  randStrSlice[5],
				slice: randStrSlice,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PosInSlice(tt.args.item, tt.args.slice); got != tt.want {
				t.Errorf("PosInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFromStringSlice(t *testing.T) {
	type args struct {
		item  string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "item3",
			args: args{
				item:  "item3",
				slice: []string{"item4", "item3", "item2", "item1", "item0"},
			},
			want: []string{"item4", "item2", "item1", "item0"},
		},
		{
			name: "rem_end-1",
			args: args{
				item:  "ksdjfh",
				slice: []string{"sadlklg", "rtkhjkjf", "skdjt54874", "ksdjfh", "dfkgj598479"},
			},
			want: []string{"sadlklg", "rtkhjkjf", "skdjt54874", "dfkgj598479"},
		},
		{
			name: "rem_first",
			args: args{
				item:  "sadlklg",
				slice: []string{"sadlklg", "rtkhjkjf", "skdjt54874", "ksdjfh", "dfkgj598479"},
			},
			want: []string{"rtkhjkjf", "skdjt54874", "ksdjfh", "dfkgj598479"},
		},
		{
			name: "rem_last",
			args: args{
				item:  "dfkgj598479",
				slice: []string{"sadlklg", "rtkhjkjf", "skdjt54874", "ksdjfh", "dfkgj598479"},
			},
			want: []string{"sadlklg", "rtkhjkjf", "skdjt54874", "ksdjfh"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFromStringSlice(tt.args.item, tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFromStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceUniq(t *testing.T) {
	type args struct {
		sslice *[]string
	}
	tests := []struct {
		name string
		args args
		want *[]string
	}{
		{
			name: "already unique",
			args: args{sslice: &[]string{"apple", "orange", "banana", "plum"}},
			want: &[]string{"apple", "orange", "banana", "plum"},
		},
		{
			name: "has duplicates",
			args: args{sslice: &[]string{"apple", "orange", "apple", "banana", "banana", "plum", "apple", "orange"}},
			want: &[]string{"apple", "orange", "banana", "plum"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceUniq(tt.args.sslice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSliceUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
