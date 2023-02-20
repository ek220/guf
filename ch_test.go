package f_test

import (
	"testing"

	"github.com/ek220/f"
	"github.com/stretchr/testify/assert"
)

func TestSliceAndCh(t *testing.T) {
	t.Parallel()

	ch := f.SliceToCh(stringSlice)
	acc := f.SliceFromCh(ch)
	assert.Equal(t, stringSlice, acc)
}

func TestFilterCh(t *testing.T) {
	t.Parallel()

	ch := f.FilterCh(f.SliceToCh(intSlice), func(v int) bool { return v%3 == 0 })
	assert.Equal(t, []int{3, -6, 9}, f.SliceFromCh(ch))
}

func TestMapCh(t *testing.T) {
	t.Parallel()

	ch := f.MapCh(f.SliceToCh(intSlice), func(v int) int { return v * 10 })
	assert.Equal(t, []int{10, -20, 30, -40, 50, -60, 70, -80, 90}, f.SliceFromCh(ch))
}

func TestSplitCh(t *testing.T) {
	t.Parallel()

	tCh, fCh := f.SplitCh(f.SliceToCh(intSlice), func(v int) bool { return v%3 == 0 })
	assert.Equal(t, []int{3, -6, 9}, f.SliceFromCh(tCh))
	assert.Equal(t, []int{1, -2, -4, 5, 7, -8}, f.SliceFromCh(fCh))
}
