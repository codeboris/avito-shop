package handler

import (
	"encoding/json"
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type AuthHandler struct {
	userService *service.UserService
	jwtSecret   string
}

func NewAuthHandler(userService *service.UserService, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "неверный запрос", http.StatusBadRequest)
		return
	}

	user, err := h.userService.AuthenticateUser(r.Context(), req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Генерация JWT-токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		http.Error(w, "ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func InitAuthHandlers(router *mux.Router, userService *service.UserService, jwtSecret string) {
	authHandler := NewAuthHandler(userService, jwtSecret)

	router.HandleFunc("/api/auth", authHandler.Auth).Methods("POST")
}
