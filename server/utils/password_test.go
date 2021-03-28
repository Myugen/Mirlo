package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	var err error
	var hashedPass string

	var pass = RandString(16)

	hashedPass, err = HashPassword(pass)
	assert.Nil(t, err)
	assert.NotNil(t, hashedPass)

	var passwordsMatch = PasswordMatch(pass, hashedPass)
	assert.True(t, passwordsMatch)
}
