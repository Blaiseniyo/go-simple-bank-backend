package models

import (
	"gorm.io/gorm"
)

type Account struct{
	gorm.Model
	Id int `json:"id" gorm:"primary_key"`
	Owner string `json:"owner"`
	Balance int `json:"balance"`
	Currency string `json:"currency"`
}