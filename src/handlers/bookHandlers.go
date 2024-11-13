package handlers

import (
	"BooksAPI/src/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBookById function returns book by id.
func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "don't use parametr ID us uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	book, ok := models.FindBookByID(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: fmt.Sprintf("book with id %v is not exists", id)}
		json.NewEncoder(writer).Encode(message)
	} else {
		writer.WriteHeader(http.StatusOK) // 200 error
		json.NewEncoder(writer).Encode(book)
	}
}

// CreateBook funtion create new book by id
func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book...")
	var book models.Book

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields() //! Отклоняет поля которые неизвестны.
	err := decoder.Decode(&book)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "provide json file is invalid."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	// Create new book and save one
	var newBookId int = len(models.DB) + 1
	book.ID = newBookId
	models.DB = append(models.DB, book)

	writer.WriteHeader(http.StatusCreated) // 401 created
	json.NewEncoder(writer).Encode(book)
}

// UpdateBookById function can update book by id.
func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating a book...")
	var userBook models.Book

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields() //! Отклоняет поля которые неизвестны.
	err := decoder.Decode(&userBook)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "provide json file is invalid."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "don't use parametr ID us uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	found := models.UpdateBookById(id, userBook)
	if !found {
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: fmt.Sprintf("book with id %v is not exists", id)}
		json.NewEncoder(writer).Encode(message)
	} else {
		updatedBook, _ := models.FindBookByID(id)
		writer.WriteHeader(http.StatusOK) // 200 error
		json.NewEncoder(writer).Encode(updatedBook)
	}
}
