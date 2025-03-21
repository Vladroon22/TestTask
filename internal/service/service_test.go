package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gthub.com/Vladroon22/TestTask/internal/entity"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateUser(c context.Context, user entity.User) error {
	args := m.Called(c, user)
	return args.Error(0)
}

func (m *MockRepo) UpdateUser(c context.Context, user entity.User) (int, error) {
	args := m.Called(c, user)
	return args.Int(0), args.Error(1)
}

func (m *MockRepo) GetUser(c context.Context, id int) (entity.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func Test_MockCreateUser(t *testing.T) {
	user := entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       23,
		Email:     "qwe@mail.ru",
		Phone:     "89115671423",
	}

	mockRepo := &MockRepo{}

	mockRepo.On("CreateUser", mock.Anything, user).Return(nil)

	service := NewService(mockRepo)

	err := service.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestService_GetUser(t *testing.T) {
	exUser := entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       23,
		Email:     "qwe@mail.ru",
		Phone:     "89115671423",
	}

	mockRepo := &MockRepo{}

	mockRepo.On("GetUser", mock.Anything, 1).Return(exUser, nil)

	service := NewService(mockRepo)

	user, err := service.GetUser(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, exUser, user)

	mockRepo.AssertExpectations(t)
}

func TestService_UpdateUser(t *testing.T) {
	user := entity.User{
		FirstName: "VVV",
		LastName:  "QQQ",
		Age:       24,
		Email:     "123@mail.ru",
		Phone:     "89117771423",
	}

	mockRepo := &MockRepo{}

	mockRepo.On("UpdateUser", mock.Anything, user).Return(1, nil)

	service := NewService(mockRepo)

	updatedID, err := service.UpdateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, 1, updatedID)

	mockRepo.AssertExpectations(t)
}
