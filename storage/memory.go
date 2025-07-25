package storage

import "github.com/effel/media-tracker/models"

var books []models.Book
var nextID = 1

func AddBook(book models.Book) models.Book {
	book.ID = nextID
	nextID++
	books = append(books, book)
	return book
}

func GetBooks() []models.Book {
	return books
}

func GetBookByID(id int) (*models.Book, bool){
	for _, book := range books{
		if book.ID == id {
			return &book, true
		}
	}
	return nil, false
}