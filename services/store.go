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
type UpdateAccountParams struct {
	Account_id int64 `json:"account_id"`
	Amount     int64 `json:"amount"`
}

type TransferResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount models.Account  `json:"from_account"`
	ToAccount   models.Account  `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

var txKey = struct{}{}

func TransferTransaction(ctx context.Context, db *gorm.DB, transfer_data TransferParams) (TransferResult, error) {
	var result TransferResult
	var err error

	db.Transaction(func(tx *gorm.DB) error {

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

		if transfer_data.From_account_id < transfer.To_account_id {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, transfer.From_account_id, -transfer_data.Amount, transfer.To_account_id, transfer_data.Amount, tx)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, transfer.To_account_id, transfer_data.Amount, transfer.From_account_id, -transfer_data.Amount, tx)
		}

		if err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return result, err
}

func addMoney(ctx context.Context, accountID1 int64, amount1 int64, accountID2 int64, amount2 int64, tx *gorm.DB) (account1 models.Account, account2 models.Account, err error) {

	account1, err = AddAccountBalance(ctx, &UpdateAccountParams{Account_id: accountID1, Amount: amount1}, tx)

	if err != nil {
		return
	}

	account2, err = AddAccountBalance(ctx, &UpdateAccountParams{Account_id: accountID2, Amount: amount2}, tx)

	return

}
