package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	baseURL      = "http://localhost:8080/users"
	sessionToken string
	isLog        = 0
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- User Management Menu ---")
		fmt.Println("1. Login")
		fmt.Println("2. Create User")
		fmt.Println("3. Get Users")
		fmt.Println("4. Update User")
		fmt.Println("5. Delete User")
		fmt.Println("6. Exit")
		fmt.Print("Select option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			login(scanner)
		case "2":
			createUser(scanner)
		case "3":
			getUsers()
		case "4":
			updateUser(scanner)
		case "5":
			deleteUser(scanner)
		case "6":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

// Авторизация
func login(scanner *bufio.Scanner) {
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	payload := map[string]string{"username": username, "password": password}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	//Если токкен подошёл
	if resp.StatusCode == http.StatusOK {
		isLog = 1
		body, _ := ioutil.ReadAll(resp.Body)
		var responseData map[string]string
		json.Unmarshal(body, &responseData)
		sessionToken = responseData["token"]
		fmt.Println("Login successful.")
	} else { //ну и если не подошёл
		isLog = 0
		fmt.Println("Failed to login. Status:", resp.Status)
	}
}

// Создание пользователя
func createUser(scanner *bufio.Scanner) {
	if isLog == 0 {
		fmt.Print("Войдите в учётную запись!!!")
		return
	}
	if isLog == 1 {
		fmt.Print("Enter name: ")
		scanner.Scan()
		name := scanner.Text()

		fmt.Print("Enter age: ")
		scanner.Scan()
		var age int
		fmt.Sscan(scanner.Text(), &age)

		fmt.Print("Enter email: ")
		scanner.Scan()
		email := scanner.Text()

		user := User{Name: name, Age: age, Email: email}
		jsonUser, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonUser))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+sessionToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			fmt.Println("User created successfully.")
		} else {
			fmt.Println("Failed to create user. Status:", resp.Status)
		}
	}
}

// Получение по id
func getUsers() {
	if isLog == 0 {
		fmt.Print("Войдите в учётную запись!!!")
		return
	}
	req, _ := http.NewRequest("GET", baseURL, nil)
	req.Header.Set("Authorization", "Bearer "+sessionToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		var users []User
		json.Unmarshal(body, &users)
		for _, user := range users {
			fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s\n", user.ID, user.Name, user.Age, user.Email)
		}
	} else {
		fmt.Println("Failed to get users. Status:", resp.Status)
	}
}

// Обновить инфу
func updateUser(scanner *bufio.Scanner) {
	if isLog == 0 {
		fmt.Print("Войдите в учётную запись!!!")
		return
	}
	fmt.Print("Enter user ID to update: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter new name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter new age: ")
	scanner.Scan()
	var age int
	fmt.Sscan(scanner.Text(), &age)

	fmt.Print("Enter new email: ")
	scanner.Scan()
	email := scanner.Text()

	user := User{Name: name, Age: age, Email: email}
	jsonUser, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", baseURL+"/"+id, bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User updated successfully.")
	} else {
		fmt.Println("Failed to update user. Status:", resp.Status)
	}
}

// Удалить
func deleteUser(scanner *bufio.Scanner) {
	if isLog == 0 {
		fmt.Print("Войдите в учётную запись!!!")
		return
	}
	fmt.Print("Enter user ID to delete: ")
	scanner.Scan()
	id := scanner.Text()

	req, _ := http.NewRequest("DELETE", baseURL+"/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+sessionToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User deleted successfully.")
	} else {
		fmt.Println("Failed to delete user. Status:", resp.Status)
	}
}
