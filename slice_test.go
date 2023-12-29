package guf_test

import (
	"testing"

	"github.com/ek220/guf"
	"github.com/stretchr/testify/assert"
)

var (
	stringSlice = []string{"a", "bc", "def", "foo", "hello"}
	intSlice    = []int{1, -2, 3, -4, 5, -6, 7, -8, 9}
	structSlice = []ts{
		{P: nil, I: 1, S: "foo"},
		{P: guf.Ptr[uint64](13), I: 4, S: "bar"},
	}
)

func TestCloneSlice(t *testing.T) {
	t.Parallel()

	x := []int{1, 2, 3, 4}
	z := guf.CloneSlice(x)
	assert.Equal(t, x, z)
	x[2] = 123
	assert.NotEqual(t, x, z)
}

func TestFilterSlice(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		[]string{"def", "foo"},
		guf.FilterSlice(stringSlice, func(s string) bool { return len(s) == 3 }))
	assert.Equal(
		t,
		[]int{-2, -4, -6, -8},
		guf.FilterSlice(intSlice, func(v int) bool { return v < 0 }))
	assert.Equal(
		t,
		[]ts{{P: guf.Ptr[uint64](13), I: 4, S: "bar"}},
		guf.FilterSlice(structSlice, func(s ts) bool { return s.P != nil }))
}

func TestMapSlice(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		[]int{1, 2, 3, 3, 5},
		guf.MapSlice(stringSlice, func(s string) int { return len(s) }))
	assert.Equal(
		t,
		[]int{11, 8, 13, 6, 15, 4, 17, 2, 19},
		guf.MapSlice(intSlice, func(v int) int { return v + 10 }))
	assert.Equal(
		t,
		[]int32{2, 8},
		guf.MapSlice(structSlice, func(s ts) int32 {
			return s.I * 2
		}))
}

func TestSplitSlice(t *testing.T) {
	t.Parallel()

	iT, iF := guf.SplitSlice(intSlice, func(v int) bool { return v < 0 })
	assert.Equal(t, []int{-2, -4, -6, -8}, iT)
	assert.Equal(t, []int{1, 3, 5, 7, 9}, iF)

	sT, sF := guf.SplitSlice(stringSlice, func(v string) bool { return len(v) != 3 })
	assert.Equal(t, []string{"a", "bc", "hello"}, sT)
	assert.Equal(t, []string{"def", "foo"}, sF)
}

func TestReverseSlice(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{5, 4, 3, 2, 1}, guf.ReverseSlice([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{}, guf.ReverseSlice([]int{}))
}
