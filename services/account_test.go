package services

import (
	"context"
	"testing"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/stretchr/testify/require"
	
)

func TestCreateAccount(t *testing.T){
	arg := models.Account{
		Owner: "Test owner",
		Balance: 100,
		Currency: "FRW",
	}

	account,err := CreateAccount(context.Background(),&arg,TEST_DB)

	require.NoError(t,err)
	require.NotEmpty(t,account)

	require.Equal(t,arg.Owner,account.Owner)
	require.Equal(t,arg.Balance,account.Balance)
	require.Equal(t,arg.Currency,account.Currency)

	require.NotZero(t,account.Currency)
	require.NotZero(t,account.CreatedAt)

}