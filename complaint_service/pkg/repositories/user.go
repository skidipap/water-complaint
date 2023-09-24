package repositories

import (
	"example/complaint_service/pkg/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func CreateUser(user *models.User) error {
	err := db.Create(user).Error
	return err
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := db.First(&user, userID).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetRoleIDByName(roleName string) (uint, error) {
	var role models.Role
	err := db.Where("name = ?", roleName).First(&role).Error
	if err != nil {
		return 5, err
	}
	return role.ID, nil
}

func DeleteUser(userID uint) error {
	var user models.User

	err := db.First(&user, userID).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		return err
	}

	err = db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
