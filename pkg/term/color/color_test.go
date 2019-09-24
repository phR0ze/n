package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyan(t *testing.T) {
	assert.Equal(t, "foobar", Cyan("foobar"))
	assert.Equal(t, "foobar", CyanB("foobar"))
	assert.Equal(t, "foobarblah", CyanL("foobar%s", "blah"))
	assert.Equal(t, "foobar", CyanBL("foobar"))

	gIsTTy = true
	assert.Equal(t, "\x1b[36mfoobar\x1b[0m", Cyan("foobar"))
	assert.Equal(t, "\x1b[1;36mfoobar\x1b[0m", CyanB("foobar"))
	assert.Equal(t, "\x1b[96mfoobarblah\x1b[0m", CyanL("foobar%s", "blah"))
	assert.Equal(t, "\x1b[1;96mfoobar\x1b[0m", CyanBL("foobar"))
}
