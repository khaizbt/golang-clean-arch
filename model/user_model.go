package model

import (
	"time"
)

type (
	User struct {
		ID        string    `json:"id" gorm:"primary_key"`
		Name      string    `json:"name" gorm:"varchar(255);NOT NULL"`
		Email     string    `json:"email" gorm:"varchar(255);NOT NULL"`
		Password  string    `json:"password" gorm:"varchar(255);NOT NULL"`
		Username  string    `json:"username" gorm:"varchar(20);NOT NULL"`
		Roles     string    `json:"roles" gorm:"varchar(10);NOT NULL"`
		Address   string    `json:"address" gorm:"varchar(255);NULL"`
		Phone     string    `json:"phone" gorm:"type:varchar(15);NULL"`
		Avatar    string    `json:"avatar" gorm:"varchar(255);NULL"`
		Status    string    `json:"status" gorm:"varchar(255);NOT NULL"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
