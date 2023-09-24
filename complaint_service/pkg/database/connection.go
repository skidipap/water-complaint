package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return database, nil

}
