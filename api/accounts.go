package api

import (
	"net/http"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/Blaiseniyo/go-simple-bank-backend/services"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EURO RFW"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := services.CreateAccount(ctx, &models.Account{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
	return
}

type AccountIDParmsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {

	var req AccountIDParmsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := services.GetAccountById(ctx, req.ID, server.db)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
	return
}

func (server *Server) deleteAccount(ctx *gin.Context) {

	var req AccountIDParmsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := services.DeleteAccount(ctx, req.ID, server.db)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
	return
}

type listAccountRequest struct {
	PageNumber int32 `form:"page_number" binding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=1,max=5"`
}

func (server *Server) listAccounts(ctx *gin.Context) {

	var req listAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset := (req.PageNumber - 1) * req.PageSize
	accounts, err := services.ListAllAccounts(ctx, int(req.PageSize), int(offset), server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
	return
}

type UpdateAccountRequest struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency" binding:"oneof=USD EURO RFW"`
	Balance  int64 `json:"balance"`
}

func (server *Server) updateAccount(ctx *gin.Context) {

	var req AccountIDParmsRequest
	var body UpdateAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	if err:= ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := services.GetAccountById(ctx, req.ID, server.db)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	updatedAccountinfo := models.Account{Owner: body.Owner,Currency: body.Currency,Balance: body.Balance}
	updatedAccount, err := services.UpdateAccount(ctx, &account,&updatedAccountinfo, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedAccount)
	return
}
