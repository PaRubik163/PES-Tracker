package middleware

import (
	"net/http"
	"strings"
	"time"
	jwt "tracker/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		status := c.Writer.Status()

		userID, ok := c.Get("user_id")

		if !ok{
			userID = "none"
		}

		if strings.HasPrefix(c.Request.URL.Path, "/api"){
			entry := logrus.WithFields(logrus.Fields{
				"user_id": userID,
				"path": c.Request.URL.Path,
				"method": c.Request.Method,
				"status": status,
				"time": start,
				"latency": time.Since(start),
			})

			switch {
			case status >= 500:
				entry.Error("Internal server error")
			case status >= 400:
				entry.Warn("Client error")
			default:
				entry.Info("Request successed")
			}
			return
		}

		if !strings.HasPrefix(c.Request.URL.Path, "/static"){
			entry := logrus.WithFields(logrus.Fields{
				"path": c.Request.URL.Path,
				"method": c.Request.Method,
				"status": c.Writer.Status(),
				"time": start,
				"latency": time.Since(start),
				"error": c.Errors.String(),
			})
			entry.Info("Frontend Request")
		}
	}
}