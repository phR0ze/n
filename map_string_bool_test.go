package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMapBool_Set(t *testing.T) {
	m := NewStringMapBool()
	assert.Equal(t, true, m.Set("test", true))
	assert.Equal(t, false, m.Set("test", true))
}
