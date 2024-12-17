package route

import (
	"goshop/config"
	"goshop/controller"
	"goshop/middleware"
	"goshop/workflow"

	"github.com/gin-gonic/gin"
)

func RouteUser(route *gin.Engine, service workflow.UserService) {
	authService := config.NewServiceAuth()
	userController := controller.NewUserController(service, authService)
	userMiddleware := middleware.AuthMiddlewareUser(authService, service) // middl.AuthMiddlewareManager(authService, workflow)
	api := route.Group("/api/v1/")
	api.POST("user/login", userController.Login)
	api.POST("update-account", userMiddleware, userController.UpdateProfile)
}
