package handlers

import (
	"Books/Packages/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	//Read Dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//Read to request body
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Description = updatedBook.Description

	h.DB.Save(&book)

	// //Iterate through all the mock books (or DB if so)
	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		book.Title = updatedBook.Title
	// 		book.Author = updatedBook.Author
	// 		book.Description = updatedBook.Description
	// 		mocks.Books[index] = book
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Header().Add("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode("Updated")
	// 		break
	// 	}
	// }

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")
}
