package utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerateUniqueID(t *testing.T) {
	uniqueID := GenerateUniqueID()

	_, err := uuid.Parse(uniqueID)
	if err != nil {
		t.Errorf("Generated ID is not a valid UUID: %v", err)
	}
}
