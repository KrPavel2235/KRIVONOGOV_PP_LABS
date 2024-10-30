package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

// Генерация и сохранение пары RSA-ключей
func generateKeys() error {
	// Генерация приватного ключа
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Сохранение приватного ключа в файл
	privateFile, err := os.Create("private_key.pem")
	if err != nil {
		return err
	}
	defer privateFile.Close()

	privatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	privateFile.Write(privatePEM)

	// Извлечение и сохранение открытого ключа
	publicKey := &privateKey.PublicKey
	publicFile, err := os.Create("public_key.pem")
	if err != nil {
		return err
	}
	defer publicFile.Close()

	publicPEM, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	pem.Encode(publicFile, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicPEM,
	})

	fmt.Println("Keys have been generated and saved to 'private_key.pem' and 'public_key.pem'")
	return nil
}

func genKey() {
	err := generateKeys()
	if err != nil {
		log.Fatalf("Error generating keys: %v", err)
	}
}

func main() {
	genKey()
}
