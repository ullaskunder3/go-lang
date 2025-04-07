package task

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const specialChars = "!@#$%^&*()-_+=<>?"

func generatePassword1(length int, useUpper, useNumbers, useSpecial bool) (string, error) {
	charset := lowercase
	if useUpper {
		charset += uppercase
	}
	if useNumbers {
		charset += numbers
	}
	if useSpecial {
		charset += specialChars
	}

	if len(charset) == 0 {
		return "", fmt.Errorf("no character sets selected")
	}

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

func getUserInputWithReader(prompt string, reader *bufio.Reader) bool {
	fmt.Printf("%s (y/n): ", prompt)
	input, _ := reader.ReadString('\n')
	return len(input) > 0 && (input[0] == 'y' || input[0] == 'Y')
}

func AdvanceCliPass() {
	fmt.Print("Enter password length: ")
	var length int
	_, err := fmt.Scan(&length)
	if err != nil || length <= 0 {
		fmt.Println("Invalid length. Please provide a positive integer.")
		os.Exit(1)
	}

	// 🧼 Flush leftover newline from fmt.Scan
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	useUpper := getUserInputWithReader("Include uppercase letters?", reader)
	useNumbers := getUserInputWithReader("Include numbers?", reader)
	useSpecial := getUserInputWithReader("Include special characters?", reader)

	password, err := generatePassword1(length, useUpper, useNumbers, useSpecial)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Generated Password:", password)
}
