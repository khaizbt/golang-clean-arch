package service

import (
	"errors"
	"goshop/input"
	"goshop/model"
	"goshop/repository"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Login(input input.LoginEmailInput) (model.User, error)
		GetUserById(ID int) (model.User, error)
	}

	service struct {
		repository repository.UserRepository
	}
)

func NewUserService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) GetUserById(ID int) (model.User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID != ID {
		return user, errors.New("User Tidak Ditemukan")
	}

	return user, nil
}

func (s *service) Login(input input.LoginEmailInput) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	// if user.StatusRegistered == 0 { //Handle if not verified
	// 	return user, errors.New("Verified Your Account and Try again")
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}
