package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// Загрузка приватного ключа из файла
func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	privatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privatePEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Неудалось расщифровать RSA ключ приватный")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// Загрузка открытого ключа из файла
func loadPublicKey(filename string) (*rsa.PublicKey, error) {
	publicPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicPEM)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("Неудалось расщифровать RSA ключ публичный")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return publicKey, nil
}

// Подписание сообщения
func signMessage(privateKey *rsa.PrivateKey, message string) (string, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, 0, hashed[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// Проверка подписи
func verifySignature(publicKey *rsa.PublicKey, message, signature string) error {
	hashed := sha256.Sum256([]byte(message))
	decodedSig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(publicKey, 0, hashed[:], decodedSig)
}

func main() {
	// Пример сообщения
	message := "Это сообщение передалось каким-то сложным образом, используя два ключа"

	// Загрузка ключей
	privateKey, err := loadPrivateKey("private_key.pem")
	if err != nil {
		log.Fatalf("Приватный ключь, эмм ну крч хз чёт не то: %v", err)
	}
	publicKey, err := loadPublicKey("public_key.pem")
	if err != nil {
		log.Fatalf("Неудалось загрузить публичный ключь: %v", err)
	}

	// Подпись сообщения
	signature, err := signMessage(privateKey, message)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	fmt.Printf("Сигнатура: %s\n", signature)

	// Проверка подписи
	err = verifySignature(publicKey, message, signature)
	if err != nil {
		fmt.Println("Верификация не удалась")
	} else {
		fmt.Println("Верификация удалась")
	}
}
