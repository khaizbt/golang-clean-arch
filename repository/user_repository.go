package repository

import (
	"github.com/khaizbt/golang-clean-arch/config"
	"github.com/khaizbt/golang-clean-arch/model"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindUserByEmail(email string) (model.User, error)
		FindByID(ID string) (model.User, error)
		UpdateProfile(user model.User) (model.User, error)
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

func (r *repository) FindByID(ID string) (model.User, error) {
	var user model.User

	err := r.db.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateProfile(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
