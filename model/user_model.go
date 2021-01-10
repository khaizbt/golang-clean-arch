package model

import (
	"time"
)

type (
	User struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Username  string    `json:"username"`
		Roles     string    `json:"roles"`
		Address   string    `json:"address"`
		Phone     int       `json:"phone"`
		Avatar    string    `json:"avatar"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
