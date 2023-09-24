package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string `json:"full_name"`
	UserName       string `json:"user_name"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"password"`
	AvatarFileName string `json:"avatar_file_name"`
	Role           Role   `json:"role"`
	RoleID         uint
}

type Role struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
