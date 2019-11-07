package bcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	var i uint64
	i = 12345678
	b := Encode(i)
	assert.Equal(t, []byte{0x78, 0x56, 0x34, 0x12}, b, "")
}
