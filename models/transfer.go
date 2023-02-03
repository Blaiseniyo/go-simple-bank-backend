package models

import (
	"gorm.io/gorm"
)

type Transfer struct{
	gorm.Model
	Id int `json:"id" gorm:"primary_key"`
	From_account_id int `json:"from_account_id"`
	To_account_id int `json:"to_account_id"`
	Amount int `json:"amount"`
}