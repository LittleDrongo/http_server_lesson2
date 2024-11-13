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

func GetBookById(writer http.ResponseWriter, request *http.Request) {
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
