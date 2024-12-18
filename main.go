package main

import (
	"fmt"
	"github.com/khaizbt/golang-clean-arch/middleware"
	"github.com/khaizbt/golang-clean-arch/repository"
	"github.com/khaizbt/golang-clean-arch/route"
	"github.com/khaizbt/golang-clean-arch/workflow"
	"log"
	"os"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	// fmt.Println("masuk", os.Getenv("DB_USER"))
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_API"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	userRepo := repository.NewUserRepository()
	userService := workflow.NewUserService(userRepo)

	secureMiddleware := middleware.SecureMiddleware()

	router := gin.Default()
	router.Use(secureMiddleware)
	router.Use(sentrygin.New(sentrygin.Options{}))
	route.RouteUser(router, userService)
	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
