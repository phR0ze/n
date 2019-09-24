package term

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTTY(t *testing.T) {
	assert.False(t, IsTTY())
}

func TestIsTTYP(t *testing.T) {
	assert.False(t, IsTTYP(os.Stdout.Fd()))
}
