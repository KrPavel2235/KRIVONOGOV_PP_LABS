package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Структура User
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// Application содержит массив пользователей вместо базы данных
type Application struct {
	Users []User
}

// Обработчик для GET /users - получение списка всех пользователей
func (app *Application) getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(app.Users)
}

// Обработчик для GET /users/{id} - получение информации о конкретном пользователе
func (app *Application) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	for _, user := range app.Users {
		if int(user.ID) == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

// Обработчик для POST /users - добавление нового пользователя
func (app *Application) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверные данные пользователя", http.StatusBadRequest)
		return
	}

	user.ID = uint(len(app.Users) + 1)
	app.Users = append(app.Users, user)
	json.NewEncoder(w).Encode(user)
}

// Обработчик для PUT /users/{id} - обновление информации о пользователе
func (app *Application) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	for i, user := range app.Users {
		if int(user.ID) == id {
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, "Неверные данные пользователя", http.StatusBadRequest)
				return
			}
			app.Users[i] = user
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

// Обработчик для DELETE /users/{id} - удаление пользователя
func (app *Application) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	for i, user := range app.Users {
		if int(user.ID) == id {
			app.Users = append(app.Users[:i], app.Users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

// Настройка маршрутов
func (app *Application) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // Middleware для логирования запросов

	r.Get("/users", app.getUsers)
	r.Get("/users/{id}", app.getUser)
	r.Post("/users", app.createUser)
	r.Put("/users/{id}", app.updateUser)
	r.Delete("/users/{id}", app.deleteUser)

	return r
}

func main() {
	app := &Application{
		Users: []User{}, // Инициализация пустого списка пользователей
	}
	http.ListenAndServe(":8080", app.routes())
}
