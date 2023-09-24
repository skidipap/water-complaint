package handlers

import (
	"example/complaint_service/pkg/models"
	"example/complaint_service/pkg/repositories"
	"example/complaint_service/pkg/usecase"
	"example/complaint_service/pkg/utils"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubmitComplaintRequest struct {
	UserID              uint                    `form:"user_id"`
	ComplaintCategoryID uint                    `form:"complaint_category_id" binding:"required"`
	MeteranID           uint                    `form:"meteran_id" binding:"required"`
	ComplaintName       string                  `form:"complaint_name" binding:"required"`
	ShortDescription    string                  `form:"short_description" binding:"required"`
	PriorityLevel       uint                    `form:"priority_level" binding:"required"`
	Images              []*multipart.FileHeader `form:"images"`
}

type UpdateComplaintRequest struct {
	ComplaintName    string `json:"complaint_name"`
	ShortDescription string `json:"short_description"`
	PriorityLevel    uint   `json:"priority_level"`
	MeteranID        uint   `json:"meteran_id"`
	CategoryID       uint   `json:"complaint_category_id"`
}

func SubmitComplaintHandler(c *gin.Context) {
	var request SubmitComplaintRequest

	// Get token for authorization
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
	// Try to bind JSON first
	if err := c.ShouldBind(&request); err != nil {
		// If binding as JSON fails, try binding as form data
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
	}

	// Retrieve Meteran and ComplaintCategory entities based on provided IDs
	complaintID, err := usecase.SubmitComplaint(
		userID, request.ComplaintCategoryID, request.MeteranID,
		request.ComplaintName, request.ShortDescription, request.PriorityLevel,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit complaint"})
		return
	}

	// // Handle file uploads
	for _, image := range request.Images {
		avatarFileName := fmt.Sprintf("complaint_image_%s%s", utils.GenerateUniqueID(), filepath.Ext(image.Filename))
		err = c.SaveUploadedFile(image, fmt.Sprintf("uploads/complaint_images/%s", avatarFileName))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save complaint image"})
			return
		}

		complaintImage := models.ComplaintImage{
			ComplaintIssueID: complaintID,
			FileName:         avatarFileName,
			IsPrimary:        false,
		}

		err = repositories.CreateComplaintImage(&complaintImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create complaint image"})
			return
		}
	}

	// Return a success response if everything is processed correctly
	c.JSON(http.StatusOK, gin.H{"message": "Complaint submitted successfully"})
}

func ListAllComplaintsHandler(c *gin.Context) {
	complaints, err := usecase.GetAllComplaintsWithImage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve complaints"})
		return
	}

	var ComplaintResponses []gin.H
	for _, complaint := range complaints {
		var imageFilenames []string
		for _, image := range complaint.ComplaintImages {
			imageFilenames = append(imageFilenames, image.FileName)
		}

		ComplaintResponse := gin.H{
			"id":                 complaint.ID,
			"user_id":            complaint.UserID,
			"complaint_category": complaint.ComplaintCategoryID,
			"complaint_name":     complaint.ComplaintName,
			"complaint_images":   imageFilenames,
			"meteran_id":         complaint.MeteranID,
			"priority_level":     complaint.PriorityLevel,
		}
		ComplaintResponses = append(ComplaintResponses, ComplaintResponse)
	}

	c.JSON(http.StatusOK, ComplaintResponses)
}

func DeleteComplaintHandler(c *gin.Context) {
	// Extract complaint ID from URL parameter
	complaintIDStr := c.Param("id")
	complaintID, err := strconv.ParseUint(complaintIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid complaint ID"})
		return
	}

	err = usecase.DeleteComplaint(uint(complaintID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete complaint"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complaint deleted successfully"})
}

func UpdateComplaintHandler(c *gin.Context) {
	// Extract complaint ID from URL parameter
	complaintIDStr := c.Param("id")
	complaintID, err := strconv.ParseUint(complaintIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid complaint ID"})
		return
	}

	var request UpdateComplaintRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err = usecase.UpdateComplaint(uint(complaintID), request.ComplaintName, request.ShortDescription, request.PriorityLevel, request.MeteranID, request.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update complaint"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complaint updated successfully"})
}
