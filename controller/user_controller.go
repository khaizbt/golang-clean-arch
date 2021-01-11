package controller

import (
	"goshop/config"
	"goshop/helper"
	"goshop/input"
	"goshop/model"
	"goshop/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
	authService config.AuthService
}

func NewUserController(userService service.UserService, authService config.AuthService) *userController {
	return &userController{userService, authService}
}

type UserFormatter struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	Phone  int    `json:"phone"`
	Token  string `json:"token"`
}

func FormatUser(user model.User, token string) UserFormatter { //Token akan didapatkan dari JWT
	formater := UserFormatter{
		UserID: user.ID,
		Email:  user.Email,
		Phone:  user.Phone,
		Token:  token,
	}

	return formater
}

func (h *userController) Login(c *gin.Context) {
	var input input.LoginEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Login Failed #LOG001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		responsError := helper.APIResponse("Login Failed #LOG002", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	token, err := h.authService.GenerateTokenUser(loggedInUser.ID)
	if err != nil {
		responsError := helper.APIResponse("Login Failed", http.StatusBadGateway, "fail", "Unable to generate token")
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	response := helper.APIResponse("Login Success", http.StatusOK, "success", FormatUser(loggedInUser, token))

	c.JSON(http.StatusOK, response)
}
