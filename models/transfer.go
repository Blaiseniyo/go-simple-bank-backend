package models

import (
	"gorm.io/gorm"
)

type Transfer struct{
	gorm.Model
	Id int64 `json:"id" gorm:"primary_key"`
	From_account_id int64 `json:"from_account_id"`
	To_account_id int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}