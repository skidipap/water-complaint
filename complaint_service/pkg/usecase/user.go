package usecase

import (
	"errors"
	"example/complaint_service/pkg/models"
	"example/complaint_service/pkg/repositories"
	"example/complaint_service/pkg/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func RegisterUser(fullname, username, password, email, avatarFileName, roleName string) error {

	roleID, err := repositories.GetRoleIDByName(roleName)
	if err != nil {
		return err
	}

	user := models.User{
		FullName:       fullname,
		UserName:       username,
		Password:       password,
		Email:          email,
		AvatarFileName: avatarFileName,
		RoleID:         roleID,
	}

	// Call the repository to create the user
	err = repositories.CreateUser(&user)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(userID uint) (*models.User, error) {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ValidateUserCredentials(email, password string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = utils.ComparePasswords(user.Password, password)
	if err != nil {
		return nil, nil
	}

	return user, nil
}

func GenerateToken(userID uint, jwtSecret []byte) (string, error) {
	// Create the claims for the JWT token
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (24 hours from now)
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractUserIdFromToken(token string, jwtSecret []byte) (uint, error) {
	claims := jwt.MapClaims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !parsedToken.Valid {
		return 0, err
	}

	// Extract the user ID from the claims
	userID := uint(claims["id"].(float64))

	return userID, nil
}

func GetUserEmailFromToken(tokenString string, jwtSecret []byte) (string, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Check if the token is valid
	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", errors.New("failed to extract claims from token")
		}

		userEmail, ok := claims["email"].(string)
		if !ok {
			return "", errors.New("failed to extract user email from token claims")
		}

		return userEmail, nil
	} else {
		return "", errors.New("invalid token")
	}
}

func GetAllUsers() ([]*models.User, error) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
