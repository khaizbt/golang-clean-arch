package route

import (
	"github.com/khaizbt/golang-clean-arch/config"
	"github.com/khaizbt/golang-clean-arch/controller"
	"github.com/khaizbt/golang-clean-arch/middleware"
	"github.com/khaizbt/golang-clean-arch/workflow"

	"github.com/gin-gonic/gin"
)

func RouteUser(route *gin.Engine, service workflow.UserService) {
	authService := config.NewServiceAuth()
	userController := controller.NewUserController(service, authService)
	userMiddleware := middleware.AuthMiddlewareUser(authService, service) // middl.AuthMiddlewareManager(authService, workflow)
	root := route.Group("/")
	root.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := root.Group("/api/v1/")
	api.POST("user/login", userController.Login)
	api.POST("update-account", userMiddleware, userController.UpdateProfile)
}
