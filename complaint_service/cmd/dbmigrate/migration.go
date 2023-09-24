package main

import (
	"example/complaint_service/pkg/database"
	"example/complaint_service/pkg/models"
	"log"
)

func main() {
	database, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&models.Role{}, &models.User{})
	database.AutoMigrate(&models.ComplaintIssue{}, &models.ComplaintCategory{}, &models.ComplaintImage{})
	database.AutoMigrate(&models.Meteran{})

}
