package models

import (
	"github.com/Trishank-K/BookStore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("Id=?", id).Delete(&book)
	return book
}
