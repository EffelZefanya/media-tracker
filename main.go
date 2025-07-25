package main

import (
	"fmt"
	"net/http"

	"github.com/effel/media-tracker/handlers"
)

func main() {
	fmt.Printf("Server is starting now...")
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/books", handlers.BooksHandler)
	http.ListenAndServe(":8080", nil)
}
