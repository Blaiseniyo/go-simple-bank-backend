package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, user *models.User, DB *gorm.DB) (models.User, error) {
	var created_user models.User
	result := DB.Create(&user).Scan(&created_user)
	return created_user, result.Error
}

func GetUser(ctx context.Context, userName string, DB *gorm.DB) (models.User, error) {
	var retrived_user models.User
	result := DB.Where("username=?",userName).First(&retrived_user)
	return retrived_user, result.Error
}
