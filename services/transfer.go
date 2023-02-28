package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateTransfer(ctx context.Context, transfer *models.Transfer, DB *gorm.DB) (models.Transfer, error) {
	var created_tranfer models.Transfer
	result := DB.Create(&transfer).Scan(&created_tranfer)
	return created_tranfer, result.Error
}

func GetTransferById(ctx context.Context, TransferId int64, DB *gorm.DB) (models.Transfer, error) {
	var transfer models.Transfer
	result := DB.First(&transfer, TransferId)
	return transfer, result.Error
}

func UpdateTransfer(ctx context.Context, transfer *models.Transfer, UpdatedTransferData *models.Transfer, DB *gorm.DB) (models.Transfer, error) {
	var updated_account models.Transfer
	result := DB.Model(&transfer).Updates(&UpdatedTransferData).Scan(&updated_account)
	return updated_account, result.Error
}

func DeleteTransfer(ctx context.Context, TransferId int64, DB *gorm.DB) (int64, error) {
	result := DB.Delete(&models.Transfer{}, TransferId)
	return result.RowsAffected, result.Error
}

func ListAllTransfers(ctx context.Context, limit int, offset int, DB *gorm.DB) ([]models.Transfer, error) {
	var transfers []models.Transfer
	result := DB.Limit(limit).Offset(offset).Find(&transfers)
	return transfers, result.Error
}
