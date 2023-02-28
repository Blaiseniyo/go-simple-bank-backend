package db

import (
	"fmt"

	"github.com/Blaiseniyo/go-simple-bank-backend/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *util.Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",config.DBUser,config.DBPassword,config.DBHost,config.DBPort,config.DBName)))
	if err != nil {
		panic(err)
	}
	return db
}
