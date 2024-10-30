package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var jwtKey = []byte("your_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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
	JWTKey   []byte
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
	r.HandleFunc("/users/login", app.login)

	return r
}

// Обработка защищенного маршрута
func (app *Application) handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected route. Only authorized users can see this message."))
}

// Метод для логина и генерации токена
func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Пример проверки пользователя
	if creds.Username != "testuser" || creds.Password != "testpassword" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// Middleware для аутентификации
func (app *Application) authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" || len(tokenString) < 7 {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		// Удаление "Bearer " из строки токена
		tokenString = tokenString[7:]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Вызов следующего обработчика
		next(w, r)
	}
}

// Основная функция запуска сервера
func main() {
	app := Application{
		Validate: validator.New(),
		JWTKey:   []byte("secret"),
	}
	app.initDB()
	logrus.Info("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", app.routes()))
}
