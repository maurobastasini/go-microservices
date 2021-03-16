package model

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_GetUser_OK(t *testing.T) {
	user, err := UserDAO.GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Mauro", user.FirstName)
	assert.EqualValues(t, "Bastasini", user.LastName)
	assert.EqualValues(t, "maurobastasiniprof@gmail.com", user.Email)
}

func Test_GetUser_NotFound(t *testing.T) {
	user, err := UserDAO.GetUser(0)

	assert.NotNil(t, err)
	assert.EqualValues(t, "userID: 0 not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Status)
	assert.Nil(t, user)
}
