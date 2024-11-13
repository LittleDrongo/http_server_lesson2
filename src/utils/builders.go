package utils

import (
	// "BooksAPI/src/handlers"

	"BooksAPI/src/handlers"

	"github.com/gorilla/mux"
)

// Отдельные функции для роутинга.
func BuildBookResource(router *mux.Router, prefix string) {

	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc(prefix, handlers.GreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookById).Methods("GET")
}

func BuildManyBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"s", handlers.GetAllBooks).Methods("GET")
}
