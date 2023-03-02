package services

import (
	"context"
	"testing"
	"time"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/Blaiseniyo/go-simple-bank-backend/util"
	"github.com/stretchr/testify/require"
)

func Createuser(t *testing.T) models.User {
	password,err := util.HashPassword(util.RandomString(6))

	require.NoError(t,err)

	arg := models.User{
		Username:    util.RandomOwner(),
		FullName:    util.RandomOwner(),
		HashedPassword:  password,
		Email: util.RandomEmail(),
	}

	user, err := CreateUser(context.Background(), &arg, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	Createuser(t)
}

func TestGet(t *testing.T) {

	new_user := Createuser(t)
	user, err := GetUser(context.Background(), new_user.Username, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, new_user.FullName, user.FullName)
	require.Equal(t, new_user.Username, user.Username)
	require.Equal(t, new_user.Email, user.Email)
	require.Equal(t, new_user.Id, user.Id)
	require.Equal(t, new_user.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.CreatedAt)
	require.WithinDuration(t, user.CreatedAt,new_user.CreatedAt,time.Second)
}