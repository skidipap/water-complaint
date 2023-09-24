package usecase

import (
	"errors"
	"example/complaint_service/pkg/models"
	"example/complaint_service/pkg/repositories"
)

type UpdateComplaintRequest struct {
	ComplaintName    string `json:"complaint_name"`
	ShortDescription string `json:"short_description"`
	PriorityLevel    uint   `json:"priority_level"`
	MeteranID        uint   `json:"meteran_id"`
	CategoryID       uint   `json:"complaint_category_id"`
}

func SubmitComplaint(userID, categoryID, meteranID uint, complaintName, shortDescription string, priorityLevel uint) (uint, error) {

	meteran, complaintCategory, err := validateAndGetEntities(meteranID, categoryID)
	if err != nil {
		return 0, err
	}

	complaint := models.ComplaintIssue{
		UserID:              userID,
		ComplaintCategoryID: categoryID,
		MeteranID:           meteranID,
		ComplaintName:       complaintName,
		ShortDescription:    shortDescription,
		PriorityLevel:       priorityLevel,
		Meteran:             *meteran,
		ComplaintCategory:   *complaintCategory,
	}

	complaintID, err := repositories.CreateComplaint(&complaint)
	if err != nil {
		return 0, err
	}

	return complaintID, nil
}

func validateAndGetEntities(meteranID, categoryID uint) (*models.Meteran, *models.ComplaintCategory, error) {
	meteran, err := repositories.GetMeteranByID(meteranID)
	if err != nil || meteran == nil {
		return nil, nil, err
	}

	complaintCategory, err := repositories.GetComplaintCategoryByID(categoryID)
	if err != nil || complaintCategory == nil {
		return nil, nil, err
	}

	return meteran, complaintCategory, nil
}

func GetAllComplaints() ([]*models.ComplaintIssue, error) {
	complaints, err := repositories.GetAllComplaints()
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetAllComplaintsWithImage() ([]*models.ComplaintIssue, error) {
	complaints, err := repositories.GetAllComplaints()
	if err != nil {
		return nil, err
	}

	for _, complaint := range complaints {
		images, err := repositories.GetAllComplaintImagesByComplaintID(complaint.ID)
		if err != nil {
			return nil, err
		}
		complaint.ComplaintImages = images
	}

	return complaints, nil
}

func DeleteComplaint(complaintID uint) error {
	err := repositories.DeleteComplaint(complaintID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateComplaint(complaintID uint, complaintName string, shortDescription string, priorityLevel uint, meteranID uint, complaintCategoryID uint) error {
	complaint, err := repositories.GetComplaintByID(complaintID)
	if err != nil {
		return err
	}

	// Check if the complaint exists
	if complaint == nil {
		return errors.New("complaint not found")
	}

	// Update the complaint properties
	complaint.ComplaintName = complaintName
	complaint.ShortDescription = shortDescription
	complaint.PriorityLevel = priorityLevel
	complaint.MeteranID = meteranID
	complaint.ComplaintCategoryID = complaintCategoryID

	err = repositories.UpdateComplaint(complaint)
	if err != nil {
		return err
	}

	return nil
}
