package user

import (
	"testing"

	"github.com/mujahxd/api3-jwt/user/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	mockRepo := new(mocks.Repository)
	service := NewService(mockRepo)

	phone := "081234567890"
	password := "passwordku"
	// data user
	user := User{

		Name:     "John Doe",
		Phone:    phone,
		Password: "$2a$10$zIjUYvn..C/UZcVyB2x4XeIXUZOy6MkT6U7T0TjlT6G7NLli1ZKKu",
	}

	mockRepo.On("FindByPhone", phone).Return(user, nil)

	// eksekusi fungsi
	result, err := service.Login(LoginInput{Phone: phone, Password: password})

	// memastikan ekspektasi terpenuhi dan mock repository tidak memiliki ekspektasi yang terlewati
	mockRepo.AssertExpectations(t)

	// memastikan hasil eksekusi fungsi sesuai dengan ekspektasi
	assert.Nil(t, err)
	assert.Equal(t, user, result)
}
