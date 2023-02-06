package controllers

import (
	"book-store-api/pkg/models"
	"book-store-api/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
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
		fmt.Printf("Error parsing bookId: %v\n", bookId)
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error parsing bookId: %v\n", bookId)
	}

	bookDetail, db := models.GetBookById(ID)
	if updatedBook.Name != "" {
		bookDetail.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		bookDetail.Author = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		bookDetail.Publication = updatedBook.Publication
	}

	db.Save(&bookDetail)

	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error parsing bookId: %v\n", bookId)
	}
	book := models.DeleteBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
