package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/golang-clean-arch/config"
	"github.com/khaizbt/golang-clean-arch/entity"
	"github.com/khaizbt/golang-clean-arch/repository"
	"github.com/khaizbt/golang-clean-arch/workflow"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var userRepo = repository.NewUserRepository()
var userService = workflow.NewUserService(userRepo)
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
