package middleware

import (
	"go_ecommerce/services"
	"go_ecommerce/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	client_token := c.Request.Header.Get("Authorization")
	if client_token == "" {
		c.JSON(http.StatusUnauthorized, utilities.UserResponse{Status: http.StatusUnauthorized, Message: "No Authorization header sent. Login and send a token!"})
		c.Abort()
		return
	}
	// client_token = client_token[7:]
	claims, err := services.ValidateToken(client_token[7:])

	if err != nil {
		c.JSON(http.StatusUnauthorized, utilities.UserResponse{Status: http.StatusUnauthorized, Message: "Invalid token!"})
		c.Abort()
		return
	}

	c.Set("user_id", claims.ID)
	c.Set("email", claims.Email)
	c.Next()
}
