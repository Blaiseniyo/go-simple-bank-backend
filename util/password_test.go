package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)


func TestPasswordHash(t *testing.T){

	password := RandomString(6)
	hashedPassword,err := HashPassword(password)
	hashedPassword2,err := HashPassword(password)

	require.NotEqual(t,hashedPassword,hashedPassword2)
	require.NoError(t,err)
	require.NotEmpty(t,hashedPassword)

	checkedPassword := CheckPassword(password,hashedPassword)

	require.NoError(t,checkedPassword)

	checkedPassword = CheckPassword(RandomString(6),hashedPassword)

	require.EqualError(t,checkedPassword,bcrypt.ErrMismatchedHashAndPassword.Error())

}