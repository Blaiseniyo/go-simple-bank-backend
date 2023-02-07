package services

import (
	"context"
	"testing"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/Blaiseniyo/go-simple-bank-backend/util"
	"github.com/stretchr/testify/require"
)

func CreateAccounts(t *testing.T) models.Account {
	arg := models.Account{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := CreateAccount(context.Background(), &arg, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.Currency)
	require.NotZero(t, account.CreatedAt)

	return account

}

func TestCreateAccount(t *testing.T) {
	CreateAccounts(t)
}

func TestGetAccount(t *testing.T) {
	
	new_account := CreateAccounts(t)
	account, err := GetAccountById(context.Background(), new_account.Id, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, new_account.Owner, account.Owner)
	require.Equal(t, new_account.Balance, account.Balance)
	require.Equal(t, new_account.Currency, account.Currency)
	require.Equal(t, new_account.Id, account.Id)

	require.NotZero(t, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	update_account := models.Account{Owner: "test_User", Balance: 12323, Currency: "FRW"}
	account := CreateAccounts(t)
	updated_account, err := UpdateAccount(context.Background(), &account, &update_account, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, updated_account.Owner)
	require.Equal(t, account.Balance, updated_account.Balance)
	require.Equal(t, account.Currency, updated_account.Currency)

	require.NotZero(t, account.UpdatedAt)
}

func TestDeletAccount(t *testing.T) {

	account := CreateAccounts(t)
	deleted_row, err := DeleteAccount(context.Background(), account.Id, TEST_DB)

	require.NoError(t, err)

	require.Equal(t, deleted_row, int64(1))
}

func TestListAllAccount(t *testing.T) {
	
	for i := 0; i < 10; i++ {
		CreateAccounts(t)
	}

	accounts, err := ListAllAccounts(context.Background(), 5, 5, TEST_DB)

	require.NoError(t, err)
	require.Equal(t, len(accounts), 5)

	for _, i := range accounts {
		require.NotEmpty(t, i)
	}
}
