package services

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TEST_DB *gorm.DB

func TestMain(m *testing.M) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable"))

	if err != nil {
		log.Fatal("can't connect to the testing database")
	}
	TEST_DB = db
	os.Exit(m.Run())
}
