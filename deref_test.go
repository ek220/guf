package f_test

import (
	"testing"

	"github.com/ek220/f"
	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	ptr := f.Ptr(13)
	assert.NotNil(t, ptr)
	assert.Equal(t, 13, *ptr)
}

func TestDerefOrDefault(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", f.DerefOrDefault[string](nil))
	assert.Equal(t, "foo", f.DerefOrDefault(f.Ptr("foo")))

	z := f.DerefOrDefault[ts](nil)
	assert.Nil(t, z.P)
	assert.Equal(t, int32(0), z.I)
	assert.Equal(t, "", z.S)
}

func TestDerefOr(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "def", f.DerefOr(nil, "def"))
	assert.Equal(t, "foo", f.DerefOr(f.Ptr("foo"), "def"))

	z := f.DerefOr(nil, ts{P: f.Ptr[uint64](13), I: 4, S: "bar"})
	assert.NotNil(t, z.P)
	assert.Equal(t, uint64(13), *z.P)
	assert.Equal(t, int32(4), z.I)
	assert.Equal(t, "bar", z.S)

}
