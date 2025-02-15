package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitHandlers(router *mux.Router) {
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userID := "user123"
		log.Printf("User not found: %v", userID)

		w.Write([]byte("OK"))
	}).Methods("GET")
}
