package services

import (
	"context"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"gorm.io/gorm"
)

type TransferParams struct {
	From_account_id int64 `json:"from_account_id"`
	To_account_id   int64 `json:"to_account_id"`
	Amount          int64 `json:"amount"`
}

type TransferResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount int64           `json:"from_account"`
	ToAccount   int64           `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

func TransferTransaction(ctx context.Context, db *gorm.DB, transfer_data *TransferParams) {
	var result TransferResult

	db.Transaction(func(tx *gorm.DB) error {

		var err error
		transfer := models.Transfer{From_account_id: transfer_data.From_account_id, To_account_id: transfer_data.To_account_id, Amount: transfer_data.Amount}
		result.Transfer, err = CreateTransfer(ctx, &transfer, tx)

		if err != nil {
			// return any error will rollback
			return err
		}

		from_entry := models.Entry{Account_id: transfer_data.From_account_id, Amount: -transfer_data.Amount}
		result.FromEntry, err = CreateEntry(ctx, &from_entry, tx)

		if err != nil {
			// return any error will rollback
			return err
		}

		to_entry := models.Entry{Account_id: transfer_data.To_account_id, Amount: transfer_data.Amount}
		result.ToEntry, err = CreateEntry(ctx, &to_entry, tx)

		if err != nil {
			// return any error will rollback
			return err
		}

		// TODO: update accounts' balance

		// return nil will commit the whole transaction
		return nil
	})
}
