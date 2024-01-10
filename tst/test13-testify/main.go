package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Database interface {
	GetBookByID(bookID string) (Book, bool)
}

type InMemoryDB struct {
	books map[string]Book
}

func (db *InMemoryDB) GetBookByID(bookID string) (Book, bool) {
	book, ok := db.books[bookID]
	return book, ok
}

func GetBookByID(w http.ResponseWriter, r *http.Request, db Database) {
	bookID := "123" // Example book ID
	book, ok := db.GetBookByID(bookID)
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Convert the book to JSON
	bookJSON, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error encoding book details", http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the book JSON to the response
	w.Write(bookJSON)
}

func main() {
	db := &InMemoryDB{
		books: map[string]Book{
			"123": {
				ID:    "123",
				Title: "Sample",
			},
		},
	}

	http.HandleFunc("/books/123", func(w http.ResponseWriter, r *http.Request) {
		GetBookByID(w, r, db)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
