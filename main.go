package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Author Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice (A length of an array variable) Book struct
var books []Book

// Header Container
func headerDRYclean(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// Get All Books
func getBooks(w http.ResponseWriter, router *http.Request) {
	headerDRYclean(w)
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, router *http.Request) {
	headerDRYclean(w)
	params := mux.Vars(router) //Get params
	// Loop through books and find it with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create A New Book
func createBook(w http.ResponseWriter, router *http.Request) {
	headerDRYclean(w)
	var book Book
	_ = json.NewDecoder(router.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID -  (not safe for production)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update The Book
func updateBook(w http.ResponseWriter, router *http.Request) {
	headerDRYclean(w)
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(router.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete The Book
func deleteBook(w http.ResponseWriter, router *http.Request) {
	headerDRYclean(w)
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init routers
	router := mux.NewRouter()
	// Mock Data - @todo - implement DB soon..
	books = append(books, Book{ID: "1", Isbn: "44873", Title: "Book One", Author: &Author{Firstname: "Fidel", Lastname: "Ramadhan"}})
	books = append(books, Book{ID: "2", Isbn: "44874", Title: "Book Two", Author: &Author{Firstname: "Revian", Lastname: "Adzani"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))

}
