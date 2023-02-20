package guf_test

import (
	"testing"

	"github.com/ek220/guf"
	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	ptr := guf.Ptr(13)
	assert.NotNil(t, ptr)
	assert.Equal(t, 13, *ptr)
}

func TestDerefOrDefault(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", guf.DerefOrDefault[string](nil))
	assert.Equal(t, "foo", guf.DerefOrDefault(guf.Ptr("foo")))

	z := guf.DerefOrDefault[ts](nil)
	assert.Nil(t, z.P)
	assert.Equal(t, int32(0), z.I)
	assert.Equal(t, "", z.S)
}

func TestDerefOr(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "def", guf.DerefOr(nil, "def"))
	assert.Equal(t, "foo", guf.DerefOr(guf.Ptr("foo"), "def"))

	z := guf.DerefOr(nil, ts{P: guf.Ptr[uint64](13), I: 4, S: "bar"})
	assert.NotNil(t, z.P)
	assert.Equal(t, uint64(13), *z.P)
	assert.Equal(t, int32(4), z.I)
	assert.Equal(t, "bar", z.S)

}
