package main

import (
	"example/complaint_service/pkg/api"
	"example/complaint_service/pkg/database"
	"example/complaint_service/pkg/repositories"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	repositories.SetDB(db)
	router := api.NewRouter()
	router.Run(":8000")
}
