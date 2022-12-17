package bookController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testproject/pkg/config"
	"testproject/pkg/models"
)

var db = config.GetDB()

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.Type", "application/json")
	var book models.Book
	book.ID = 1234
	book.Author = "Amir Alam"
	book.Name = "Horror House"
	book.Publication = "dexk labs"
	// json.NewDecoder(r.Body).Decode(&book)
	db.Create(&book)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
