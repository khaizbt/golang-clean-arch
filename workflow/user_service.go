package workflow

import (
	"errors"
	"github.com/khaizbt/golang-clean-arch/entity"
	"github.com/khaizbt/golang-clean-arch/model"
	"github.com/khaizbt/golang-clean-arch/repository"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Login(input entity.LoginEmailInput) (model.User, error)
		GetUserById(ID string) (model.User, error)
		UpdateProfile(input entity.DataUserInput) (bool, error)
	}

	service struct {
		repository repository.UserRepository
	}
)

func NewUserService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) GetUserById(ID string) (model.User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID != "" {
		return user, errors.New("User Tidak Ditemukan")
	}

	return user, nil
}

func (s *service) Login(input entity.LoginEmailInput) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
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

func (s *service) UpdateProfile(input entity.DataUserInput) (bool, error) {
	user, err := s.repository.FindByID(input.ID)

	if err != nil {
		return false, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Username = input.Username

	_, err = s.repository.UpdateProfile(user)

	if err != nil {
		return false, err
	}

	return true, nil
}
