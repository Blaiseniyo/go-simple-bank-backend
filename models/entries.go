package models

import (
	"gorm.io/gorm"
)

type Entry struct{
	gorm.Model
	Id int `json:"id" gorm:"primary_key"`
	Account_id int `json:"account_id"`
	Amount int `json:"amount"`
}