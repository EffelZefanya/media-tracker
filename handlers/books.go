package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/effel/media-tracker/models"
	"github.com/effel/media-tracker/storage"
	"github.com/effel/media-tracker/utils"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetBooks(w, r)
	case http.MethodPost:
		handlePostBooks(w, r)
	default:
		http.Error(w, "Method not allowed.\n", http.StatusMethodNotAllowed)
	}
}

func BooksByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Expected path: /books/{id}
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetBooksByID(w, r, id)
	case http.MethodDelete:
		handleDeleteBooksByID(w, r, id)
	default:
		http.Error(w, "Method not allowed.\n", http.StatusMethodNotAllowed)
	}
}


func handlePostBooks(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	errMsg := utils.ValidateBook(newBook)
	if errMsg != "" {
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	createdBook := storage.AddBook(newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

func handleGetBooks(w http.ResponseWriter, r *http.Request) {
	books := storage.GetBooks()

	jsonData, err := json.Marshal(books)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleDeleteBooksByID(w http.ResponseWriter, r *http.Request, id int) {
	deleted := storage.DeleteBookByID(id)
	if !deleted {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}


func handleGetBooksByID(w http.ResponseWriter, r *http.Request, id int) {
	book, found := storage.GetBookByID(id)
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

