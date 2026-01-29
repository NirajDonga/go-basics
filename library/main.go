package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func booksHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		json.NewEncoder(w).Encode(books)

	case "POST":
		var book Book
		err := json.NewDecoder(req.Body).Decode(&book)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		}

		books = append(books, book)

	case "DELETE":
		var bookID string
		err := json.NewDecoder(req.Body).Decode(&bookID)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		i := 0
		for _, b := range books {
			if b.ID != bookID {
				books[i] = b
				i++
			}
		}
		books = books[:i]

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/books", booksHandler)
	http.ListenAndServe(":8080", nil)
}
