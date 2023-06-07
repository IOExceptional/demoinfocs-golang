package bitread_test

import (
	"bytes"
	"testing"

	bit "github.com/markus-wa/demoinfocs-golang/v3/internal/bitread"
	"github.com/stretchr/testify/assert"
)
func TestReadBits(t *testing.T) {
	t.Parallel()

	var (
		buf []byte
		r   *bit.BitReader
	)

	buf = []byte{0x10, 0x01, 0x00, 0x1a, 0x0f}
	r = bit.NewSmallBitReader(bytes.NewReader(buf))
	
	assert.Equal(t, uint(16), r.ReadInt(5))
	assert.Equal(t, uint(8), r.ReadInt(5))
	assert.Equal(t, uint(0), r.ReadInt(5))
	assert.Equal(t, uint(0), r.ReadInt(5))
	assert.Equal(t, uint(0), r.ReadInt(5))
	assert.Equal(t, uint(13), r.ReadInt(5))
	assert.Equal(t, uint(28), r.ReadInt(5))
	assert.Equal(t, uint(1), r.ReadInt(5))
}