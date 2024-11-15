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

	_, found := models.FindBookByID(id)
	if !found {
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: fmt.Sprintf("book with id %v is not exists", id)}
		json.NewEncoder(writer).Encode(message)
		return
	}

	ok := models.UpdateBookById(id, userBook)
	if !ok {
		writer.WriteHeader(http.StatusNoContent) // 204 error
		return
	} else {
		writer.WriteHeader(http.StatusOK) // 200 error
		message := models.Message{Message: fmt.Sprintf("book with id %v has updated", id)}
		json.NewEncoder(writer).Encode(message)
	}
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	_, ok := models.FindBookByID(id)

	if !ok {
		log.Println("book with id =", id, "not found.")
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	// Delete book
	models.DeleteBookById(id)
	message := models.Message{Message: "book has successfully deleted from database."}
	json.NewEncoder(writer).Encode(message)
}
