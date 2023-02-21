package services

import (
	"context"
	"fmt"
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
		txName := fmt.Sprintf("tx %d",i+1)
		go func() {
			ctx := context.WithValue(context.Background(),txKey,txName)
			result, err := TransferTransaction(ctx, TEST_DB, TransferParams{
				From_account_id: account1.Id,
				To_account_id:   account2.Id,
				Amount:          amount,
			})
			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)
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

		// checkout
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.Id, account1.Id)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.Id, toAccount.Id)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)
		k := int(diff1 / amount)
		require.NotContains(t, existed, k)
		require.True(t, k >= 1 && k <= n)
		existed[k] = true
	}

	updatedFromAccount, err := GetAccountById(context.Background(), account1.Id, TEST_DB)
	require.NoError(t, err)

	updatedToAccount, err := GetAccountById(context.Background(), account2.Id, TEST_DB)
	require.NoError(t, err)
	fmt.Println(">> After:",updatedFromAccount.Balance, updatedToAccount.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedFromAccount.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedToAccount.Balance)
}
