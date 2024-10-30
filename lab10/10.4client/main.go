package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Загрузка клиентского сертификата и ключа
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatalf("Ошибка загрузки клиентского сертификата: %v", err)
	}

	// Загрузка корневого сертификата (CA) для проверки сертификата сервера
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("Ошибка загрузки корневого сертификата: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Настройка TLS-конфигурации для клиента
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
	}

	// Подключение к серверу
	conn, err := tls.Dial("tcp", "localhost:8080", tlsConfig)
	if err != nil {
		log.Fatalf("Ошибка подключения к серверу: %v", err)
	}
	defer conn.Close()

	fmt.Println("Соединение с сервером установлено")

	// Чтение и отправка сообщений
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите сообщение: ")
		message, _ := reader.ReadString('\n')

		// Отправка сообщения серверу
		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Fatalf("Ошибка отправки сообщения: %v", err)
		}

		// Получение ответа от сервера
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Ошибка получения ответа: %v", err)
		}
		fmt.Print("Ответ от сервера: ", response)
	}
}
