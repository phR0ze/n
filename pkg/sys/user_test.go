package sys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Home(t *testing.T) {
	user, err := CurrentUser()
	assert.Nil(t, err)
	home, err := UserHome()
	assert.Nil(t, err)

	assert.Equal(t, home, user.Home)
}
