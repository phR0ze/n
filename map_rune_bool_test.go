package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneMapBool_Set(t *testing.T) {
	m := NewRuneMapBool()
	assert.Equal(t, true, m.Set('a', true))
	assert.Equal(t, false, m.Set('a', true))
}
