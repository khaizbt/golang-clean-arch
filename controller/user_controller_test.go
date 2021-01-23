package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"goshop/config"
	"goshop/entity"
	"goshop/repository"
	"goshop/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var userRepo = repository.NewUserRepository()
var userService = service.NewUserService(userRepo)
var authService = config.NewServiceAuth()
var userTestController = NewUserController(userService, authService)

func TestUserController_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	input := entity.LoginEmailInput{
		Email:    "khaiz@ggggg.com",
		Password: "larashop",
	}
	requestBody, _ := json.Marshal(input)
	r := gin.Default()
	r.POST("/api/v1/login", userTestController.Login)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	//r.Run(":8000", w)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
