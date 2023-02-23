package main

import (
	"log"

	"github.com/Blaiseniyo/go-simple-bank-backend/api"
	"github.com/Blaiseniyo/go-simple-bank-backend/db"
	"github.com/Blaiseniyo/go-simple-bank-backend/util"
	"gorm.io/gorm"
)

// import (
//
//	"net/http"
//
// )
var DB *gorm.DB


func main() {

	config,err := util.LoadConfig(".")
	DB = db.Connect(&config)
	server := api.NewServer(DB)
	err = server.Start(config.ServeAddress)
	if err != nil {
		log.Fatal(err)
	}
}
