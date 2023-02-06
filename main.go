package main

import (
	// "context"
	// "fmt"
	// "log"

	"github.com/Blaiseniyo/go-simple-bank-backend/db"
	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/Blaiseniyo/go-simple-bank-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// import (
//   "net/http"
// )
var DB *gorm.DB
func main() {
  // r := gin.Default()
  // r.GET("/ping", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{
  //     "message": "pong",
  //   })
  // })
  // r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  
  DB = db.Connect()
  // accounts := []models.Account{} 
  account := models.Account{
    Owner: "Peaple Account",
    Currency: "FRW",
    Balance: 100000,
    } 
  c := &gin.Context{}
  services.CreateAccount(c,&account,DB)
  // accounts,err := services.ListAllAccounts(c,DB)
  // if err != nil{
  //   log.Fatal("can not retries data from the database")
  // }
  // DB.Create(&account)
  // DB.Find(&accounts)
  // fmt.Println(accounts)
}