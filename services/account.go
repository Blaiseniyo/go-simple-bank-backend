package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateAccount(ctx context.Context, account *models.Account, DB *gorm.DB) (models.Account, error) {
	var created_account models.Account
	result := DB.Create(&account).Scan(&created_account)
	return created_account, result.Error
}

func GetAccountById(ctx context.Context, accountId int64, DB *gorm.DB) (models.Account, error) {
	var retrived_account models.Account
	result := DB.First(&retrived_account, accountId)
	return retrived_account, result.Error
}

// func GetAccountByIdForUpdate(ctx context.Context, accountId int64, DB *gorm.DB) ( models.Account, error) {
// 	var retrived_account models.Account
// 	result := DB.Set("gorm:query_option", "FOR UPDATE").First(&retrived_account, accountId)
// 	// result := DB.First(&retrived_account, accountId)
// 	return retrived_account, result.Error
// }

func UpdateAccount(ctx context.Context, account *models.Account, UpdatedAccountData *models.Account, DB *gorm.DB) (models.Account, error) {
	result := DB.Model(&account).Updates(&UpdatedAccountData)
	return *account, result.Error
}

func AddAccountBalance(ctx context.Context, UpdatedAccountData *UpdateAccountParams, DB *gorm.DB) (models.Account, error) {
	account := models.Account{}
	result := DB.Raw("UPDATE accounts SET balance = balance + ? WHERE id = ? RETURNING *", &UpdatedAccountData.Amount, &UpdatedAccountData.Account_id).Scan(&account)
	return account, result.Error
}

func DeleteAccount(ctx context.Context, accountId int64, DB *gorm.DB) (int64, error) {
	result := DB.Delete(&models.Account{}, accountId)
	return result.RowsAffected, result.Error
}

func ListAllAccounts(ctx context.Context, limit int, offset int, DB *gorm.DB) ([]models.Account, error) {
	var accounts []models.Account
	result := DB.Limit(limit).Offset(offset).Find(&accounts)
	return accounts, result.Error
}
