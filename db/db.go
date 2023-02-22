package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable"))
	if err != nil {
		panic(err)
	}
	return db
}
