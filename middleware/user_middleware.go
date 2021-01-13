package middleware

import (
	"goshop/config"
	"goshop/helper"
	"goshop/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func AuthMiddlewareUser(authService config.AuthService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") { //Cek apakah di auth Header ada kata Bearer atau tidak
			response := helper.APIResponse("Unauthorized #TKN001", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ") //Memisahkan Bearer dengan Token dan kembalian Array(Bearer[0], Token[1])

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1] //Mengambil index ke 1 dari array token
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {

			response := helper.APIResponse("Unauthorized #TKN002", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized #TKN003", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		userID := int(claim["user_id"].(float64)) //Claim type datanya MapClaim, harus diubah ke int(sesuai parameter GetUserByID)
		role := claim["role"]

		user, err := userService.GetUserById(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized #TKN004", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //Agar proses dihentikan/tidak eksekusi program yang dibungkus middleware
			return
		}

		c.Set("currentUser", user) //Menyimpan current user ke c.CurrentUser
		c.Set("role", role)

	}

}
