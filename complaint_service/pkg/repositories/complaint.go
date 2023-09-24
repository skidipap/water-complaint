package repositories

import (
	"example/complaint_service/pkg/models"

	"gorm.io/gorm"
)

func GetComplaintCategoryByID(id uint) (*models.ComplaintCategory, error) {
	var category models.ComplaintCategory
	err := db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateComplaint(complaint *models.ComplaintIssue) (uint, error) {
	return complaint.ID, db.Create(complaint).Error
}

func CreateComplaintImage(image *models.ComplaintImage) error {
	if err := db.Create(image).Error; err != nil {
		return err
	}
	return nil
}

func GetAllComplaints() ([]*models.ComplaintIssue, error) {
	var complaints []*models.ComplaintIssue

	if err := db.Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetAllComplaintImagesByComplaintID(complaintID uint) ([]*models.ComplaintImage, error) {
	var images []*models.ComplaintImage

	err := db.Where("complaint_issue_id = ?", complaintID).Find(&images).Error
	if err != nil {
		return nil, err
	}

	return images, nil
}

func DeleteComplaint(complaintID uint) error {
	err := db.Delete(&models.ComplaintIssue{}, complaintID).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateComplaint(complaint *models.ComplaintIssue) error {
	err := db.Save(complaint).Error
	if err != nil {
		return err
	}

	return nil
}

func GetComplaintByID(complaintID uint) (*models.ComplaintIssue, error) {
	var complaint models.ComplaintIssue
	err := db.First(&complaint, complaintID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &complaint, nil
}
