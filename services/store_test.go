package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {

	account1 := CreateAccounts(t)
	account2 := CreateAccounts(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := TransferTransaction(context.Background(), TEST_DB, TransferParams{
				From_account_id: account1.Id,
				To_account_id:   account2.Id,
				Amount:          amount,
			})
			errs <- err
			results <- result
		}()
	}

	// check results

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.Id, transfer.From_account_id)
		require.Equal(t, account2.Id, transfer.To_account_id)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.Id)
		require.NotZero(t, transfer.CreatedAt)

		_, err = GetTransferById(context.Background(), transfer.Id, TEST_DB)
		require.NoError(t, err)

		// check from entry
		FromEntry := result.FromEntry
		require.NotEmpty(t, FromEntry)
		require.Equal(t, account1.Id, FromEntry.Account_id)
		require.Equal(t, -amount, FromEntry.Amount)
		require.NotZero(t, FromEntry.Id)
		require.NotZero(t, FromEntry.CreatedAt)

		_, err = GetEntryById(context.Background(), FromEntry.Id, TEST_DB)
		require.NoError(t, err)

		// check to entry

		ToEntry := result.ToEntry
		require.NotEmpty(t, ToEntry)
		require.Equal(t, account2.Id, ToEntry.Account_id)
		require.Equal(t, amount, ToEntry.Amount)
		require.NotZero(t, ToEntry.Id)
		require.NotZero(t, ToEntry.CreatedAt)

		_, err = GetEntryById(context.Background(), ToEntry.Id, TEST_DB)
		require.NoError(t, err)

		// TODO: check account updated amounts
	}
}
