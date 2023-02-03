package models

import (
	"gorm.io/gorm"
)

type Account struct{
	gorm.Model
	Id int64 `json:"id" gorm:"primary_key"`
	Owner string `json:"owner"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
}