package repositories

import (
	"example/complaint_service/pkg/models"
)

func GetMeteranByID(id uint) (*models.Meteran, error) {
	var meteran models.Meteran
	err := db.First(&meteran, id).Error
	if err != nil {
		return nil, err
	}
	return &meteran, nil
}
