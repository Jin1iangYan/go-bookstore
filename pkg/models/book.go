package models

import (
	"github.com/Jin1iangYan/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Publication string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() Book {
	db.NewRecord(b)
	db.Create(b)
	return *b
}

func GetAllBools() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBoolById(id int64) (Book, *gorm.DB) {
	var book Book
	db := db.Where("id = ?", id).Find(&book)
	return book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(book)
	return book
}
