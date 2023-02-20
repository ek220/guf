package guf_test

import (
	"testing"

	"github.com/ek220/guf"
	"github.com/stretchr/testify/assert"
)

func TestSliceAndCh(t *testing.T) {
	t.Parallel()

	ch := guf.SliceToCh(stringSlice)
	acc := guf.SliceFromCh(ch)
	assert.Equal(t, stringSlice, acc)
}

func TestFilterCh(t *testing.T) {
	t.Parallel()

	ch := guf.FilterCh(guf.SliceToCh(intSlice), func(v int) bool { return v%3 == 0 })
	assert.Equal(t, []int{3, -6, 9}, guf.SliceFromCh(ch))
}

func TestMapCh(t *testing.T) {
	t.Parallel()

	ch := guf.MapCh(guf.SliceToCh(intSlice), func(v int) int { return v * 10 })
	assert.Equal(t, []int{10, -20, 30, -40, 50, -60, 70, -80, 90}, guf.SliceFromCh(ch))
}

func TestSplitCh(t *testing.T) {
	t.Parallel()

	tCh, fCh := guf.SplitCh(guf.SliceToCh(intSlice), func(v int) bool { return v%3 == 0 })
	assert.Equal(t, []int{3, -6, 9}, guf.SliceFromCh(tCh))
	assert.Equal(t, []int{1, -2, -4, 5, 7, -8}, guf.SliceFromCh(fCh))
}
