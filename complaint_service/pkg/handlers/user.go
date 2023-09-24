// /pkg/handlers/user_handler.go

package handlers

import (
	"example/complaint_service/pkg/repositories"
	"example/complaint_service/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	Token string `json:"token" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	RoleID   uint   `json:"role_id"`
}

func DeleteUser(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	// Extract the user ID from the token
	userID, err := usecase.ExtractUserIdFromToken(token, jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	user, err := repositories.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.ID == userID {
		err := repositories.DeleteUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
	}

}

func ListAllUsers(c *gin.Context) {
	users, err := usecase.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var userResponses []UserResponse

	for _, user := range users {
		userResponse := UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.UserName,
			FullName: user.FullName,
			RoleID:   user.RoleID,
		}
		userResponses = append(userResponses, userResponse)
	}

	c.JSON(http.StatusOK, userResponses)
}
