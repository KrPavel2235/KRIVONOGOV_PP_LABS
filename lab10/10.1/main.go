package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"
)

// Функция для выбора хэш-функции по названию
func selectHashFunction(algorithm string) (hash.Hash, error) {
	switch algorithm {
	case "md5":
		return md5.New(), nil
	case "sha256":
		return sha256.New(), nil
	case "sha512":
		return sha512.New(), nil
	default:
		return nil, fmt.Errorf("Не поддерживает алгоритм: %s", algorithm)
	}
}

// Функция для хеширования строки
func hashString(algorithm, input string) (string, error) {
	hashFunc, err := selectHashFunction(algorithm)
	if err != nil {
		return "", err
	}
	hashFunc.Write([]byte(input))
	return hex.EncodeToString(hashFunc.Sum(nil)), nil
}

// Функция для проверки целостности данных
func verifyIntegrity(algorithm, input, expectedHash string) (bool, error) {
	calculatedHash, err := hashString(algorithm, input)
	if err != nil {
		return false, err
	}
	return calculatedHash == expectedHash, nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Использование: go run main.go <algorithm> <string> [<expectedHash>]")
	}

	algorithm := os.Args[1]
	input := os.Args[2]

	// Вычисление хэша для строки
	hashValue, err := hashString(algorithm, input)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Использован %s hash: %s\n", algorithm, hashValue)

	// Проверка целостности, если предоставлен ожидаемый хэш
	if len(os.Args) > 3 {
		expectedHash := os.Args[3]
		isValid, err := verifyIntegrity(algorithm, input, expectedHash)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		if isValid {
			fmt.Println("Хэш и строка совпадают")
		} else {
			fmt.Println("Хэш и строка НЕ совпадают")
		}
	}
}
