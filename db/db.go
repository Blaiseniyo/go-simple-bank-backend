package db

import (
	// "fmt"

	// "github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB { 
	
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable"))
	if err != nil {
		panic(err)
	}
	// fmt.Println(db.Statement)
	// db.AutoMigrate(&models.Account{})
	return db

}