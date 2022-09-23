package handlers

import (
	"Books/Packages/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	//Read to request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var book models.Book
	json.Unmarshal(body, &book)

	//Append to the Book mock (or to the DB)
	//mocks.Books = append(mocks.Books, book)

	//Append to DB
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Send a 201 response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
