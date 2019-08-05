package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book is
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author is
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//init book var as a slice[array]
var books []Book

//get all book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create book
func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //MOck ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

//update the book information
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	//init router
	router := mux.NewRouter()

	//MOck data
	books = append(books, Book{ID: "1", Isbn: "8546952", Title: "Book One", Author: &Author{Firstname: "J0hn", Lastname: "Wick"}})
	books = append(books, Book{ID: "2", Isbn: "8596962", Title: "Book TWO", Author: &Author{Firstname: "J0hn", Lastname: "Wick2"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
