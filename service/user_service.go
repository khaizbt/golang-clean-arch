package service

import (
	"goshop/repository"
)

type (
	UserService interface {
	}

	service struct {
		repository repository.UserRepository
	}
)

func NewUserService(repository repository.UserRepository) *service {
	return &service{repository}
}
