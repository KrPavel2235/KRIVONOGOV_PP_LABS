package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// Функция для генерации AES блока шифрования
func createCipherBlock(key string) (cipher.Block, error) {
	// AES-256 использует 32 байта ключа
	const keySize = 32
	// Приводим ключ к нужной длине, если он короче
	if len(key) < keySize {
		key = key + string(make([]byte, keySize-len(key)))
	}
	return aes.NewCipher([]byte(key)[:keySize])
}

// Функция для шифрования строки с использованием AES
func encrypt(data, key string) (string, error) {
	block, err := createCipherBlock(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(data)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Генерируем IV (инициализационный вектор) случайным образом
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Функция для расшифрования строки с использованием AES
func decrypt(encryptedData, key string) (string, error) {
	block, err := createCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	// Проверяем размер шифрованного текста
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: go run main.go <encrypt|decrypt> <text> <key>")
	}

	action := os.Args[1]
	text := os.Args[2]
	key := os.Args[3]

	switch action {
	case "encrypt":
		encryptedText, err := encrypt(text, key)
		if err != nil {
			log.Fatalf("Error encrypting text: %v", err)
		}
		fmt.Printf("Encrypted text: %s\n", encryptedText)
	case "decrypt":
		decryptedText, err := decrypt(text, key)
		if err != nil {
			log.Fatalf("Error decrypting text: %v", err)
		}
		fmt.Printf("Decrypted text: %s\n", decryptedText)
	default:
		log.Fatal("Invalid action. Use 'encrypt' or 'decrypt'.")
	}
}
