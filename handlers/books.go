package handlers

import (
	"encoding/json"
	"net/http"

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
