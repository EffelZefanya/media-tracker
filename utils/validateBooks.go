package utils

import "github.com/effel/media-tracker/models"

func ValidateBook(book models.Book) string{
	if book.Title == ""{
		return "Title is required"
	}else if book.Author == ""{
		return "Author is required"
	}

	validStatuses:= map[string]bool{
		"reading": true,
		"completed": true,
		"planned": true,
	}

	if _, ok := validStatuses[book.Status]; !ok {
		return "Status must be 'reading', 'completed' or 'planned'"
	}
	return ""
}