package repository

import (
	"learn/go/internal/database"
	"learn/go/internal/models"
)

func CreateUser(user models.Users) (models.Users, error) {
	result := database.DB.Create(&user)
	return user, result.Error
}
