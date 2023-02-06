package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateAccount(ctx context.Context, account *models.Account, DB *gorm.DB) (* models.Account,error){
	result:= DB.Create(&account)
	return account,result.Error
}

func GetAccountById(ctx context.Context, account *models.Account,accountId int64, DB *gorm.DB) (* models.Account,error){
	result:= DB.First(&account,accountId)
	return nil,result.Error
}

func UpdateAccount(ctx context.Context, account *models.Account,UpdatedAccountData *models.Account, DB *gorm.DB) (* models.Account,error){
	result:= DB.Model(&account).Updates(&UpdatedAccountData)
	return account,result.Error
}

func DeleteAccount(ctx context.Context, accountId int64, DB *gorm.DB) (int64, error){
	result:= DB.Delete(&models.Account{}, accountId)
	return result.RowsAffected, result.Error
}

// func ListAllAccounts(ctx context.Context, DB *gorm.DB) ([] models.Account,error){
// 	accounts := []models.Account{} 
// 	result := DB.Find(&accounts)
// 	return accounts,result.Error
// }