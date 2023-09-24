package utils

import "testing"

func TestComparePasswords(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %s", err)
	}

	err = ComparePasswords(hashedPassword, password)
	if err != nil {
		t.Errorf("Error comparing passwords: %s", err)
	}

}
