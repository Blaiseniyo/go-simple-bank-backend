package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil{
		return "", fmt.Errorf("fail to hashpassword: %w", err)
	}

	return string(hashed_password),nil
}

func CheckPassword(password string,hashedpassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword),[]byte(password))
}