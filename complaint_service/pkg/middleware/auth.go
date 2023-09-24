package middleware

import (
	"example/complaint_service/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	if err := utils.ValidateToken(tokenString); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Next()
}
