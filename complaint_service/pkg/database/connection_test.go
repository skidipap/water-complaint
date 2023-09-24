package database

import (
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	// Call the ConnectDatabase function
	db, err := ConnectDatabase()

	// Check if there was an error connecting to the database
	if err != nil {
		t.Errorf("Error connecting to the database: %v", err)
	}

	// Check if the database object is not nil
	if db == nil {
		t.Error("Database object is nil")
	}
}
