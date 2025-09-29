package middleware

import (
	"net/http"
	"strings"
	jwt "tracker/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwt *jwt.Jwt) gin.HandlerFunc {
	return func(c *gin.Context){
		token := c.GetHeader("Authorization")

		if token == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
			return
		}

		if strings.HasPrefix(token, "Bearer "){
			token = strings.TrimPrefix(token, "Bearer ")
		}

		session, err := jwt.ValidateToken(token)

		if err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
			return
		}

		c.Set("uuid", session.ID)
		c.Set("user_id", session.UserID)
		c.Next()
	}
}