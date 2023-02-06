package models

import (
	"book-store-api/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Automatically call - similar to contructor
func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

// Add a function to the struct
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where(`ID=?`, id).Find(&book)
	return &book, db
}

func DeleteBookById(id int64) Book {
	var book Book
	db.Where(`ID=?`, id).Delete(&book)
	return book
}
