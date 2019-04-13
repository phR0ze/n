package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrMapBool_Set(t *testing.T) {
	m := NewStrMapBool()
	assert.Equal(t, true, m.Set(*A("test"), true))
	assert.Equal(t, false, m.Set(*A("test"), true))
}
