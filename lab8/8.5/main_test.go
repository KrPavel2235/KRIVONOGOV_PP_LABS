package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User модель данных
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"gte=0"`
	Email string `json:"email" validate:"required,email"`
}

// Application содержит базу данных и валидатор
type Application struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

// Инициализация базы данных
func (app *Application) initDB() {
	dsn := "host=localhost user=postgres password=123 dbname=mydatabase port=5432 sslmode=disable"
	var err error
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	app.DB.AutoMigrate(&User{})
}

// Создание пользователя
func (app *Application) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}
	if err := app.Validate.Struct(user); err != nil {
		http.Error(w, "Ошибка валидации данных", http.StatusBadRequest)
		return
	}
	if err := app.DB.Create(&user).Error; err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Получение списка пользователей с пагинацией и фильтрацией
func (app *Application) getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	name := r.URL.Query().Get("name")
	ageParam := r.URL.Query().Get("age")
	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("page_size")
	var page, pageSize int
	var err error

	// Параметры пагинации
	if page, err = strconv.Atoi(pageParam); err != nil || page < 1 {
		page = 1
	}
	if pageSize, err = strconv.Atoi(pageSizeParam); err != nil || pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := app.DB.Offset(offset).Limit(pageSize)
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if ageParam != "" {
		age, err := strconv.Atoi(ageParam)
		if err == nil {
			query = query.Where("age = ?", age)
		}
	}
	query.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// Получение конкретного пользователя
func (app *Application) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user User
	if err := app.DB.First(&user, id).Error; err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Обновление информации о пользователе
func (app *Application) updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user User
	if err := app.DB.First(&user, id).Error; err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}
	if err := app.Validate.Struct(user); err != nil {
		http.Error(w, "Ошибка валидации данных", http.StatusBadRequest)
		return
	}
	app.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// Удаление пользователя
func (app *Application) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := app.DB.Delete(&User{}, id).Error; err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Настройка маршрутов
func (app *Application) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/users", app.getUsers)
	r.Get("/users/{id}", app.getUser)
	r.Post("/users", app.createUser)
	r.Put("/users/{id}", app.updateUser)
	r.Delete("/users/{id}", app.deleteUser)
	return r
}

func main() {
	app := Application{
		Validate: validator.New(),
	}
	app.initDB()
	logrus.Info("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", app.routes()))
}

// Функция для инициализации тестовой базы данных в памяти
func setupTestDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&User{})
	return db
}

// Функция для создания нового приложения с тестовой базой данных
func setupApp() *Application {
	app := &Application{
		DB:       setupTestDB(),
		Validate: validator.New(),
	}
	return app
}

// Тест для создания пользователя
func TestCreateUser(t *testing.T) {
	app := setupApp()
	r := chi.NewRouter()
	r.Post("/users", app.createUser)

	// Создаем данные для нового пользователя
	user := User{
		Name:  "Test User",
		Age:   30,
		Email: "testuser@example.com",
	}
	body, _ := json.Marshal(user)

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Создаем HTTP-ответ
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Проверяем статус ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидаемый статус %v, но получили %v", http.StatusOK, status)
	}

	// Проверяем тело ответа
	var createdUser User
	if err := json.NewDecoder(rr.Body).Decode(&createdUser); err != nil {
		t.Errorf("Ошибка декодирования JSON: %v", err)
	}

	if createdUser.Name != user.Name || createdUser.Email != user.Email || createdUser.Age != user.Age {
		t.Errorf("Созданный пользователь не совпадает с ожиданием")
	}
}

// Тест для получения списка пользователей
func TestGetUsers(t *testing.T) {
	app := setupApp()
	r := chi.NewRouter()
	r.Get("/users", app.getUsers)

	// Создаем HTTP-запрос
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	// Создаем HTTP-ответ
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Проверяем статус ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидаемый статус %v, но получили %v", http.StatusOK, status)
	}

	// Проверяем, что ответ содержит массив пользователей
	var users []User
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		t.Errorf("Ошибка декодирования JSON: %v", err)
	}
}
