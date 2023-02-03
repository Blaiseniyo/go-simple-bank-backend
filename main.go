package main

import (
	"fmt"

	"github.com/Blaiseniyo/go-simple-bank-backend/db"
	"github.com/Blaiseniyo/go-simple-bank-backend/models"
)

// import (
//   "net/http"
//   "github.com/gin-gonic/gin"
// )

func main() {
  // r := gin.Default()
  // r.GET("/ping", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{
  //     "message": "pong",
  //   })
  // })
  // r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  DB := db.Connect()
  accounts := []models.Account{} 
  account := models.Account{
    Owner: "Blaise niyonkuru",
    Currency: "FRW",
    Balance: 1000000,
  } 
  DB.Create(&account)
  DB.Find(&accounts)
  fmt.Println(accounts)
}