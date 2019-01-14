package n

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	files, err := filepath.Glob("test_test.go")
	assert.Nil(t, err)
	assert.Equal(t, []string{}, files)
}
