package handlers

import (
	"Books/Packages/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
