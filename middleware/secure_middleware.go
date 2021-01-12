package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func SecureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			FrameDeny:               true,
			CustomFrameOptionsValue: "SAMEORIGIN",
		})

		err := secureMiddleware.Process(c.Writer, c.Request)

		if err != nil {
			c.Abort()
			return
		}

		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	}
}
