package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Id         int64 `json:"id" gorm:"primary_key"`
	Account_id int64 `json:"account_id"`
	Amount     int64 `json:"amount"`
}
