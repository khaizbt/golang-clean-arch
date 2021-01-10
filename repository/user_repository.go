package repository

import (
	"goshop/config"
	"goshop/model"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindUserByEmail(email string) (model.User, error)
	}

	repository struct {
		db *gorm.DB
	}
)

func NewUserRepository() *repository {
	return &repository{config.GetDB()}
}

func (r *repository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
