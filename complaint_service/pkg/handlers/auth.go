package handlers

import (
	"example/complaint_service/pkg/usecase"
	"example/complaint_service/pkg/utils"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("secret123")

type RegisterRequest struct {
	Fullname   string                `json:"fullname" form:"fullname"`
	Username   string                `json:"username" form:"username"`
	Password   string                `json:"password" form:"password"`
	Email      string                `json:"email" form:"email"`
	AvatarFile *multipart.FileHeader `form:"avatar_file"`
}

func RegisterHandler(c *gin.Context) {
	var request RegisterRequest

	// Try to bind JSON first
	if err := c.ShouldBind(&request); err != nil {
		// If binding as JSON fails, try binding as form data
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Handle file upload
	if request.AvatarFile != nil {
		// Generate a unique filename (or use a simpler approach if you prefer)
		avatarFileName := "avatar_" + filepath.Base(request.AvatarFile.Filename)

		// save avatar file
		err = c.SaveUploadedFile(request.AvatarFile, "uploads/avatars/"+avatarFileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar file"})
			return
		}

		err = usecase.RegisterUser(request.Fullname, request.Username, hashedPassword, request.Email, avatarFileName, "User")
	} else {
		err = usecase.RegisterUser(request.Fullname, request.Username, hashedPassword, request.Email, "", "User")
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func LoginHandler(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := usecase.ValidateUserCredentials(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	fmt.Println(string(jwtSecret))
	// use userID for token generation
	token, err := usecase.GenerateToken(user.ID, jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
