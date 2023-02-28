package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{db: db}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
