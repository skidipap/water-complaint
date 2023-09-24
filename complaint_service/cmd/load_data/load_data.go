package main

import (
	"example/complaint_service/pkg/database"
	"example/complaint_service/pkg/models"
	"fmt"
	"log"
	"math/rand"

	"gorm.io/gorm"
)

func GenerateRoles(db *gorm.DB) {

	roles := []models.Role{
		{Name: "Surveyor", Description: "Administrator role"},
		{Name: "Maintenance", Description: "Administrator role"},
		{Name: "Customer Service", Description: "Administrator role"},
		{Name: "Engineering", Description: "Administrator role"},
		{Name: "User", Description: "Regular User"},
	}

	db.Create(&roles)

}

func GenerateUser(db *gorm.DB) {
	var surveyorRole models.Role
	var maintenanceRole models.Role
	var customerServiceRole models.Role
	var engineeringRole models.Role
	var userRole models.Role

	db.First(&surveyorRole, "name = ?", "Surveyor")
	db.First(&maintenanceRole, "name = ?", "Maintenance")
	db.First(&customerServiceRole, "name = ?", "Customer Service")
	db.First(&engineeringRole, "name = ?", "Engineering")
	db.First(&userRole, "name = ?", "User")

	// generate admin / officer user
	users := []models.User{
		{
			FullName:       "Surveyor User",
			UserName:       "surveyoruser",
			Email:          "surveyor@example.com",
			Password:       "surveyorpassword",
			AvatarFileName: "surveyor_avatar.jpg",
			Role:           surveyorRole,
		},
		{
			FullName:       "Maintenance User",
			UserName:       "maintenanceuser",
			Email:          "maintenance@example.com",
			Password:       "maintenancepassword",
			AvatarFileName: "maintenance_avatar.jpg",
			Role:           maintenanceRole,
		},
		{
			FullName:       "Customer Service User",
			UserName:       "customerserviceuser",
			Email:          "customerservice@example.com",
			Password:       "customerservicepassword",
			AvatarFileName: "customerservice_avatar.jpg",
			Role:           customerServiceRole,
		},
		{
			FullName:       "Engineering User",
			UserName:       "engineeringuser",
			Email:          "engineering@example.com",
			Password:       "engineeringpassword",
			AvatarFileName: "engineering_avatar.jpg",
			Role:           engineeringRole,
		},
	}

	// Generate 5 regular users
	for i := 0; i < 5; i++ {
		user := models.User{
			FullName:       fmt.Sprintf("Regular User %d", i+1),
			UserName:       fmt.Sprintf("regularuser%d", i+1),
			Email:          fmt.Sprintf("regularuser%d@example.com", i+1),
			Password:       fmt.Sprintf("regularpassword%d", i+1),
			AvatarFileName: fmt.Sprintf("regular_avatar%d.jpg", i+1),
			Role:           userRole,
		}
		users = append(users, user)
	}

	db.Create(&users)

}

func GenerateMeteran(db *gorm.DB) {
	var regularUsers []models.User
	db.Where("role_id = ?", 5).Find(&regularUsers) // Assuming UserRole has ID 5

	meterans := []models.Meteran{}

	for i := 0; i < 10; i++ {
		// Pick a random regular user
		randomUser := regularUsers[rand.Intn(len(regularUsers))]

		// Generate random latitude and longitude
		latitude := rand.Float64()*(90.0-(-90.0)) - 90.0
		longitude := rand.Float64()*(180.0-(-180.0)) - 180.0

		meteran := models.Meteran{
			User:        randomUser,
			MeteranCode: fmt.Sprintf("REG%d", i+1),
			Address:     fmt.Sprintf("Regular User's Address %d", i+1),
			Latitude:    latitude,
			Longitude:   longitude,
		}
		meterans = append(meterans, meteran)
	}

	db.Create(&meterans)
}

func GenerateComplaintCategory(db *gorm.DB) {
	categories := []models.ComplaintCategory{
		{Name: "Leak", Class: "complaint"},
		{Name: "Water Pressure", Class: "complaint"},
		{Name: "Water Quality", Class: "complaint"},
		{Name: "Billing", Class: "complaint"},
		{Name: "Other", Class: "complaint"},
	}

	db.Create(&categories)
}

func main() {
	database, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	GenerateRoles(database)
	GenerateUser(database)
	GenerateMeteran(database)
	GenerateComplaintCategory(database)
}
