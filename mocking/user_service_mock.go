package user

import "github.com/stretchr/testify/mock"

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}
