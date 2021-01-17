package main

import (
	"crypto/tls"
	"fmt"
	"goshop/middleware"
	"goshop/repository"
	"goshop/route"
	"goshop/service"
	"log"
	"net/http"
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
	userService := service.NewUserService(userRepo)

	secureMiddleware := middleware.SecureMiddleware()

	router := gin.Default()
	router.Use(secureMiddleware)
	router.Use(sentrygin.New(sentrygin.Options{}))
	route.RouteUser(router, userService)

	// Using This For Secure Request TLS

	certPair1, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}

	tlsConfig := new(tls.Config)
	tlsConfig.NextProtos = []string{"http/1.1"}
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.PreferServerCipherSuites = true

	tlsConfig.Certificates = []tls.Certificate{
		certPair1, /** add other certificates here **/
	}
	tlsConfig.BuildNameToCertificate()

	tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven
	tlsConfig.CurvePreferences = []tls.CurveID{
		tls.CurveP521,
		tls.CurveP384,
		tls.CurveP256,
	}
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	}

	serverTls := http.Server{
		Addr:      ":8000",
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	serverTls.ListenAndServeTLS("", "")
}
