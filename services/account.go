package services

import (
	"context"
	"fmt"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

func CreateAccount(ctx context.Context, account *models.Account, DB *gorm.DB) (* models.Account,error){
	err:= DB.Create(&account)
	 fmt.Println(err)
	return account,nil
}

func ListAllAccounts(ctx context.Context, DB *gorm.DB) ([] models.Account,error){
	accounts := []models.Account{} 
	DB.Find(&accounts)
	fmt.Println(accounts)
	return accounts,nil
}