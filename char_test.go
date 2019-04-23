package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChar_String(t *testing.T) {

	// nil
	{
		var char *Char
		assert.Equal(t, "", char.String())
		assert.Equal(t, "", NewChar("").String())

		c := Char(0)
		assert.Equal(t, "", c.String())
	}

	// bytes
	{
		assert.Equal(t, "t", NewChar(byte(0x74)).String())
		assert.Equal(t, "t", NewChar([]byte{0x74}).String())
	}

	// ints
	{
		assert.Equal(t, "2", NewChar(2).String())
		assert.Equal(t, "4", NewChar(int64(4)).String())
		assert.Equal(t, 't', NewChar(uint8(0x74)).O())
	}

	// rune
	{
		assert.Equal(t, "1", NewChar('1').String())
	}

	// string
	{
		assert.Equal(t, "1", NewChar("1").String())
	}
}
