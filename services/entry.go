package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateEntry(ctx context.Context, Entry *models.Entry, DB *gorm.DB) (*models.Entry, error) {
	result := DB.Create(&Entry)
	return Entry, result.Error
}

func GetEntryById(ctx context.Context, Entry *models.Entry, EntryId int64, DB *gorm.DB) (*models.Entry, error) {
	result := DB.First(&Entry, EntryId)
	return nil, result.Error
}

func UpdateEntry(ctx context.Context, Entry *models.Entry, UpdatedEntryData *models.Entry, DB *gorm.DB) (*models.Entry, error) {
	result := DB.Model(&Entry).Updates(&UpdatedEntryData)
	return Entry, result.Error
}

func DeleteEntry(ctx context.Context, EntryId int64, DB *gorm.DB) (int64, error) {
	result := DB.Delete(&models.Entry{}, EntryId)
	return result.RowsAffected, result.Error
}

func ListAllEntries(ctx context.Context, Entrys *[]models.Entry, limit int, offset int, DB *gorm.DB) (*[]models.Entry, error) {
	result := DB.Limit(limit).Offset(offset).Find(&Entrys)
	return Entrys, result.Error
}
