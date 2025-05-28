package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"learn/go/internal/database"
	"learn/go/internal/models"
	"learn/go/internal/repository"
	"learn/go/internal/requests"
)

func Register(request requests.RegisterRequest) (models.Users, error) {

	email := request.Email
	var existingUser models.Users
	err := database.DB.Where("email = ?", email).First(&existingUser).Error

	if err == nil {
		return models.Users{}, fmt.Errorf("user with email '%s' already exists", email)
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := models.Users{
		Email:    email,
		Password: string(hashedPass),
	}
	return repository.CreateUser(user)
}
