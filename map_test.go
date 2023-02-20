package f_test

import (
	"testing"

	"github.com/ek220/f"
	"github.com/stretchr/testify/assert"
)

var (
	stringMap = map[string]int{"a": 11, "bc": 22, "def": 33, "foo": 44, "hello": 55}
	intMap    = map[int]string{1: "a", -2: "bc", 3: "asd", -4: "we", 5: "trr"}
)

func TestCloneMap(t *testing.T) {
	t.Parallel()

	x := map[string]int{"a": 1, "b": 2}
	z := f.CloneMap(x)
	assert.Equal(t, x, z)
	x["a"] = 123
	assert.NotEqual(t, x, z)
}

func TestFilterMap(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		map[string]int{"def": 33, "foo": 44},
		f.FilterMap(stringMap, func(k string, v int) bool { return len(k) == 3 }))
	assert.Equal(
		t,
		map[int]string{3: "asd", 5: "trr"},
		f.FilterMap(intMap, func(k int, v string) bool { return len(v) == 3 }))
}

func TestMapMap(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		map[string]int{"a": 1, "bc": 2, "def": 3, "foo": 3, "hello": 5},
		f.MapMap(stringMap, func(k string, v int) (string, int) { return k, len(k) }))
	assert.Equal(
		t,
		map[int]string{100: "aa", -200: "bcbc", 300: "asdasd", -400: "wewe", 500: "trrtrr"},
		f.MapMap(intMap, func(k int, v string) (int, string) { return k * 100, v + v }))
}

func TestSplitMap(t *testing.T) {
	t.Parallel()

	sT, sF := f.SplitMap(stringMap, func(k string, v int) bool { return len(k) != 3 })
	assert.Equal(t, map[string]int{"a": 11, "bc": 22, "hello": 55}, sT)
	assert.Equal(t, map[string]int{"def": 33, "foo": 44}, sF)

	iT, iF := f.SplitMap(intMap, func(k int, v string) bool { return k < 0 })
	assert.Equal(t, map[int]string{-2: "bc", -4: "we"}, iT)
	assert.Equal(t, map[int]string{1: "a", 3: "asd", 5: "trr"}, iF)
}
