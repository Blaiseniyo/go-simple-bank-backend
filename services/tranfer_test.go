package services

import (
	"context"
	"testing"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/stretchr/testify/require"
)

func createTransfer(t *testing.T) models.Transfer {

	account_1 := CreateAccounts(t)
	account_2 := CreateAccounts(t)
	transfer_data := models.Transfer{From_account_id: account_1.Id, To_account_id: account_2.Id, Amount: 1}
	transfer, err := CreateTransfer(context.Background(), &transfer_data, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer_data.From_account_id, transfer.From_account_id)
	require.Equal(t, transfer_data.To_account_id, transfer.To_account_id)
	require.Equal(t, transfer.Amount, transfer.Amount)

	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createTransfer(t)
}

func TestGetTransfer(t *testing.T) {

	new_transfer := createTransfer(t)
	transfer, err := GetTransferById(context.Background(), new_transfer.Id, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, new_transfer.From_account_id, transfer.From_account_id)
	require.Equal(t, new_transfer.To_account_id, transfer.To_account_id)
	require.Equal(t, new_transfer.Amount, transfer.Amount)
	require.Equal(t, new_transfer.Id, transfer.Id)

	require.NotZero(t, transfer.CreatedAt)

}

func TestUpdateTransfer(t *testing.T) {

	update_transfer := models.Transfer{Amount: 2}
	transfer := createTransfer(t)
	updated_transfer, err := UpdateTransfer(context.Background(), &transfer, &update_transfer, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, updated_transfer)

	require.Equal(t, updated_transfer.Amount, update_transfer.Amount)

	require.NotZero(t, updated_transfer.UpdatedAt)
}

func TestDeleteTransfer(t *testing.T) {

	transfer := createTransfer(t)
	deleted_row, err := DeleteTransfer(context.Background(), transfer.Id, TEST_DB)

	require.NoError(t, err)

	require.Equal(t, deleted_row, int64(1))
}

func TestListAllTransfers(t *testing.T) {

	for i := 0; i < 10; i++ {
		createTransfer(t)
	}

	transfers, err := ListAllTransfers(context.Background(), 5, 5, TEST_DB)

	require.NoError(t, err)
	require.Equal(t, len(transfers), 5)

	for _, i := range transfers {
		require.NotEmpty(t, i)
	}
}
