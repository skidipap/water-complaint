package database_te_test

import (
	"example/complaint_service/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase(t *testing.T) {
	db, err := database.ConnectDatabase()
	assert.NotNil(t, db, "Expected a non-nil database instance")
	assert.Nil(t, err, "Expected no error when connecting to the database")

}
