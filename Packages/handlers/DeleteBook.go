package handlers

import (
	"Books/Packages/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	//Read Dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Iterate through all the mock books (or DB if so)
	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		// mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
	// 		// w.WriteHeader(http.StatusOK)
	// 		// json.NewEncoder(w).Encode("Deleted")
	// 		break
	// 	}
	// }

	// Find the Book by ID in our DB

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Delete by ID
	h.DB.Delete(&book)

	//mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
