package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Trishank-K/BookStore/pkg/models"
	"github.com/Trishank-K/BookStore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
		json.NewEncoder(w).Encode("Error While Parsing")
		return
	}
	bookDetails, _ := models.GetBookById(ID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		fmt.Println("Error While Parsing into JSON")
		json.NewEncoder(w).Encode("Error While Parsing into JSON")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
		json.NewEncoder(w).Encode("Error While Parsing")
		return
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
		json.NewEncoder(w).Encode("Error While Parsing")
		return
	}
	bookDetails, db := models.GetBookById(bookId)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
