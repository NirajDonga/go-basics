package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type IDRequest struct {
	ID string `json:"id"`
}

var (
	books []Book
	mu    sync.RWMutex
)

func booksHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		mu.RLock()
		json.NewEncoder(w).Encode(books)
		mu.RUnlock()

	case "POST":
		var reqBody BookRequest

		err := json.NewDecoder(req.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		book := Book{
			ID:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Title:  reqBody.Title,
			Author: reqBody.Author,
		}

		mu.Lock()
		books = append(books, book)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)

	case "DELETE":
		var reqBody IDRequest

		err := json.NewDecoder(req.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		mu.Lock()

		found := false
		i := 0

		for _, b := range books {
			if b.ID == reqBody.ID {
				found = true
				continue
			}

			books[i] = b
			i++
		}

		if !found {
			mu.Unlock()
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		books = books[:i]
		mu.Unlock()

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/books", booksHandler)
	http.ListenAndServe(":8080", nil)
}
