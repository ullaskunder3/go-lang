package task

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=<>?"

func generatePassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[randIndex.Int64()]
	}
	return string(password), nil
}

func CliPasswordGen() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <password_length>")
		os.Exit(1)
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil || length <= 0 {
		fmt.Println("Invalid length. Please provide a positive integer.")
		os.Exit(1)
	}

	password, err := generatePassword(length)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Generated Password:", password)
}
