package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Соединение закрыто клиентом")
				return
			}
			fmt.Println("Ошибка чтения сообщения:", err)
			return
		}

		fmt.Println("Получено сообщение:", message)

		_, err = conn.Write([]byte("Сообщение получено!\n"))
		if err != nil {
			fmt.Println("Ошибка отправки подтверждения:", err)
		}
	}
}

func main() {
	// Загрузка сертификата сервера и ключа
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Ошибка загрузки сертификата сервера: %v", err)
	}

	// Загрузка корневого сертификата (CA) для проверки клиентского сертификата
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("Ошибка загрузки корневого сертификата: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Настройки TLS с обязательной проверкой клиентского сертификата
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	listener, err := tls.Listen("tcp", "localhost:8080", tlsConfig)
	if err != nil {
		log.Fatalf("Ошибка создания слушателя TLS: %v", err)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на localhost:8080 с поддержкой TLS")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Ошибка приема соединения: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}
