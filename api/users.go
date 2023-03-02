package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/Blaiseniyo/go-simple-bank-backend/services"
	"github.com/Blaiseniyo/go-simple-bank-backend/util"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type CreateUserResponse struct {
	Username string `json:"username" binding:"required,alpnum"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreateAt time.Time `json:"create_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userData := models.User{
		Username:       req.Username,
		FullName:       req.Fullname,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	createdUser, err := services.CreateUser(ctx, &userData, server.db)

	if err != nil {
		if strings.Contains(err.Error(), "sqlstate 23505") {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response := CreateUserResponse{
		Username: createdUser.Username,
		Fullname: createdUser.FullName,
		Email:    createdUser.Email,
		PasswordChangedAt: createdUser.PasswordChangedAt,
		CreateAt: createdUser.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
	return
}
