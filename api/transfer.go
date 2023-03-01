package api

import (
	"fmt"
	"github.com/Blaiseniyo/go-simple-bank-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transferAmountRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) transferAmount(ctx *gin.Context) {
	var req transferAmountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccout(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validAccout(ctx, req.FromAccountID, req.Currency) {
		return
	}

	result, err := services.TransferTransaction(ctx, server.db, services.TransferParams{
		From_account_id: req.FromAccountID,
		To_account_id:   req.ToAccountID,
		Amount:          req.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}

func (server *Server) validAccout(ctx *gin.Context, accountID int64, currency string) bool {

	account, err := services.GetAccountById(ctx, accountID, server.db)

	if err != nil {
		if err != nil {
			if err.Error() == "record not found" {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return false
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return false
		}
	}

	if account.Currency != currency {
		err = fmt.Errorf("account [%d] currency mismatch %s vs %s", account.Id, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
