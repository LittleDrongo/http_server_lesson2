package main

import (
	"BooksAPI/src/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/*
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
*/

const (
	apiPrefiex string = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      string = apiPrefiex + "/book"  // URL: /api/v1/book
	manyBooksResourcePrefix string = apiPrefiex + "/books" // URL :api/v1/books
)

func init() {
	err := godotenv.Load() //by default path: "./.env"
	if err != nil {
		log.Fatalf("Cound nit found .env file :%v", err)
	}

	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBookResource(router, manyBooksResourcePrefix)

	log.Println("Router initializing successfully. Ready to go.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
