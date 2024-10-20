package utils

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenValue(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteArray := make([]byte, length)
	_, err := rand.Read(byteArray)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		byteArray[i] = charset[int(byteArray[i])%len(charset)]
	}

	return string(byteArray), nil
}

func HashPassword(salt, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(salt, password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
	return err == nil
}
