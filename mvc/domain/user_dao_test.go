package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_NotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "user 0 is not expected to be retrieved")
	assert.NotNil(t, err, "error should be not found instead of nil")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user with id 0 was not found", err.Message)
}

func TestGetUser_OK(t *testing.T) {
	user, err := GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Mauro", user.FirstName)
	assert.EqualValues(t, "Bastasini", user.LastName)
	assert.EqualValues(t, "maurobastasiniprof@gmail.com", user.Email)
}