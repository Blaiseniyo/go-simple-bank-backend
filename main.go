package main

import (
	"github.com/Blaiseniyo/go-simple-bank-backend/api"
	"github.com/Blaiseniyo/go-simple-bank-backend/db"
	"gorm.io/gorm"
	"log"
)

// import (
//
//	"net/http"
//
// )
var DB *gorm.DB

const (
	address = "0.0.0.0:8080"
)

func main() {

	DB = db.Connect()
	server := api.NewServer(DB)
	err := server.Start(address)
	if err != nil {
		log.Fatal(err)
	}
}
