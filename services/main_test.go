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

	var err error
	TEST_DB, err = gorm.Open(postgres.Open("postgres://postgres:password@localhost:5433/simple_bank?sslmode=disable"))

	if err != nil {
		log.Fatal("can't connect to the testing database")
	}

	os.Exit(m.Run())
}
