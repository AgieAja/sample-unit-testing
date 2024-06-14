package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	mockUserService := new(MockUserService)
	expectedUser := User{Id: 1, Name: "John"}

	mockUserService.On("GetUser", 1).Return(expectedUser, nil)

	user, err := mockUserService.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockUserService.AssertExpectations(t)
}
