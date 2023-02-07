package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateEntry(ctx context.Context, Entry *models.Entry, DB *gorm.DB) (models.Entry, error) {
	var created_entry models.Entry
	result := DB.Create(&Entry).Scan(&created_entry)
	return created_entry, result.Error
}

func GetEntryById(ctx context.Context, EntryId int64, DB *gorm.DB) (models.Entry, error) {
	var entry models.Entry
	result := DB.First(&entry, EntryId)
	return entry, result.Error
}

func UpdateEntry(ctx context.Context, entry *models.Entry, UpdatedEntryData *models.Entry, DB *gorm.DB) (models.Entry, error) {
	result := DB.Model(&entry).Updates(&UpdatedEntryData)
	return *entry, result.Error
}

func DeleteEntry(ctx context.Context, EntryId int64, DB *gorm.DB) (int64, error) {
	result := DB.Delete(&models.Entry{}, EntryId)
	return result.RowsAffected, result.Error
}

func ListAllEntries(ctx context.Context, limit int, offset int, DB *gorm.DB) ([]models.Entry, error) {
	var entries []models.Entry
	result := DB.Limit(limit).Offset(offset).Find(&entries)
	return entries, result.Error
}
