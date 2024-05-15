// main.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct represents a book entity
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book.ID = len(books) + 1
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for index, book := range books {
		if book.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&books[index]); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			books[index].ID = id // Ensure ID remains the same
			json.NewEncoder(w).Encode(books[index])
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
