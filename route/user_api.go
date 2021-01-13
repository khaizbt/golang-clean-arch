package route

import (
	"goshop/config"
	"goshop/controller"
	"goshop/middleware"
	"goshop/service"

	"github.com/gin-gonic/gin"
)

func RouteUser(route *gin.Engine, service service.UserService) {
	authService := config.NewServiceAuth()
	userController := controller.NewUserController(service, authService)
	userMiddleware := middleware.AuthMiddlewareUser(authService, service) // middl.AuthMiddlewareManager(authService, service)
	api := route.Group("/api/v1/")
	api.POST("login", userController.Login)
	api.POST("update-account", userMiddleware, userController.UpdateProfile)
}
