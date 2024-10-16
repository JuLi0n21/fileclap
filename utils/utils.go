package utils

import "crypto/rand"

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
