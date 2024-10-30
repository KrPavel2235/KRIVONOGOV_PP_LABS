package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"` // "admin" or "user"
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func main() {
	r := chi.NewRouter()

	r.Post("/login", loginHandler)
	r.With(authMiddleware).Get("/admin", adminHandler)
	r.With(authMiddleware).Get("/user", userHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Role == "" {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Генерация токена
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	// Отправка токена клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Необходим токен", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		// Сохранение роли пользователя в контексте запроса
		r = r.WithContext(context.WithValue(r.Context(), "role", claims.Role))

		next.ServeHTTP(w, r)
	})
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	if role != "admin" {
		http.Error(w, "Доступ запрещен", http.StatusForbidden)
		return
	}
	w.Write([]byte("Добро пожаловать, администратор!"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	if role != "user" && role != "admin" {
		http.Error(w, "Доступ запрещен", http.StatusForbidden)
		return
	}
	w.Write([]byte("Добро пожаловать, пользователь!"))
}
