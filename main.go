package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
	Price  float64 `json:"price"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for index, book := range books {
		if book.ID == vars["id"] {
			books = append(books[:index], books[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
	// Initialize the books slice with some data
	books = []Book{
		{ID: "1", Title: "1984", Author: &Author{Firstname: "George", Lastname: "Orwell"}, Price: 19.84},
		{ID: "2", Title: "Brave New World", Author: &Author{Firstname: "Aldous", Lastname: "Huxley"}, Price: 18.99},
		{ID: "3", Title: "Fahrenheit 451", Author: &Author{Firstname: "Ray", Lastname: "Bradbury"}, Price: 15.99},
		{ID: "4", Title: "The Catcher in the Rye", Author: &Author{Firstname: "J.D.", Lastname: "Salinger"}, Price: 10.99},
		{ID: "5", Title: "To Kill a Mockingbird", Author: &Author{Firstname: "Harper", Lastname: "Lee"}, Price: 12.99},
	}

	// Create a new router
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start the server
	fmt.Printf("Starting server on port 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}