package user

import (
	"errors"

	"github.com/mujahxd/api3-jwt/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsPhoneAvailable(input CheckPhoneInput) (bool, error)
	GetUserByID(ID int) (User, error)
	UpdateUser(inputID int, inputData RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Phone = input.Phone
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	phone := input.Phone
	password := input.Password

	user, err := s.repository.FindByPhone(phone)
	if err != nil {
		return user, err
	}

	// cek user
	if user.ID == 0 {
		return user, errors.New("no user found on that phone")
	}

	// password validation
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) IsPhoneAvailable(input CheckPhoneInput) (bool, error) {
	phone := input.Phone

	user, err := s.repository.FindByPhone(phone)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return true, nil

}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	// cek user
	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}

	return user, nil
}

func (s *service) UpdateUser(inputID int, inputData RegisterUserInput) (User, error) {
	user, err := s.repository.FindByID(inputID)
	if err != nil {
		return user, err
	}
	user.Name = inputData.Name
	hashedPassword, err := utils.HashPassword(inputData.Password)

	if err != nil {
		return user, err
	}

	user.Password = hashedPassword

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
