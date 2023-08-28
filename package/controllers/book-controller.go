package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Apurvapingale/book-store/package/config"
	"github.com/Apurvapingale/book-store/package/models"
	"github.com/Apurvapingale/book-store/package/utils"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	//send response to user
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")                 //this to console
		w.Header().Set("Content-Type", "pkglication/json") //this to client
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Unable to parse the request"}`))
	}
	bookDetails := models.GetBookById(ID)

	//send response to user
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	err := utils.ParseBody(r, &NewBook)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := NewBook.CreateBook()

	//send response to user
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)

	//send response to user
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	db := config.GetDB()
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publisher != "" {
		bookDetails.Publisher = updateBook.Publisher
	}
	db.Save(&bookDetails)

	//send response to user
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
