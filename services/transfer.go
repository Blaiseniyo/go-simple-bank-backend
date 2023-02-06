package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateTransfer(ctx context.Context, Transfer *models.Transfer, DB *gorm.DB) (*models.Transfer, error) {
	result := DB.Create(&Transfer)
	return Transfer, result.Error
}

func GetTransferById(ctx context.Context, Transfer *models.Transfer, TransferId int64, DB *gorm.DB) (*models.Transfer, error) {
	result := DB.First(&Transfer, TransferId)
	return nil, result.Error
}

func UpdateTransfer(ctx context.Context, Transfer *models.Transfer, UpdatedTransferData *models.Transfer, DB *gorm.DB) (*models.Transfer, error) {
	result := DB.Model(&Transfer).Updates(&UpdatedTransferData)
	return Transfer, result.Error
}

func DeleteTransfer(ctx context.Context, TransferId int64, DB *gorm.DB) (int64, error) {
	result := DB.Delete(&models.Transfer{}, TransferId)
	return result.RowsAffected, result.Error
}

func ListAllTransfers(ctx context.Context, Transfers *[]models.Transfer, limit int, offset int, DB *gorm.DB) (*[]models.Transfer, error) {
	result := DB.Limit(limit).Offset(offset).Find(&Transfers)
	return Transfers, result.Error
}
