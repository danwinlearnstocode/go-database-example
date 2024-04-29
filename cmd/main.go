package main

import (
	"database/sql"
	db "example/go-database-example/pkg/database"
	"example/go-database-example/pkg/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Article REST API!")
    fmt.Println("Article REST API")
}

func handleRequests(DB *sql.DB) {
		h := handlers.New(DB)
    // create a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
		myRouter.HandleFunc("/articles", h.GetAllArticles).Methods(http.MethodGet)
		myRouter.HandleFunc("/articles/{id}", h.GetArticle).Methods(http.MethodGet)
		myRouter.HandleFunc("/articles", h.AddArticle).Methods(http.MethodPost)
		myRouter.HandleFunc("/articles/{id}", h.UpdateArticle).Methods(http.MethodPut)
		myRouter.HandleFunc("/articles/{id}", h.DeteleArticle).Methods(http.MethodDelete)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	DB := db.Connect()
	db.CreateTable(DB)
  handleRequests(DB)
	db.CloseConnection(DB)
}