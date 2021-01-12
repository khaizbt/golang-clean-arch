package route

import (
	"goshop/config"
	"goshop/controller"
	"goshop/service"

	"github.com/gin-gonic/gin"
)

func RouteUser(route *gin.Engine, secureMiddleware gin.HandlerFunc, service service.UserService) {
	authService := config.NewServiceAuth()
	userController := controller.NewUserController(service, authService)

	api := route.Group("/api/v1/")
	api.POST("login", secureMiddleware, userController.Login)
}
