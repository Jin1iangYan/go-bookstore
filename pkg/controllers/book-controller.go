package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jin1iangYan/go-bookstore/pkg/models"
	"github.com/Jin1iangYan/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

// This is not the book model,
// think of it as a serializer
type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var NewBook models.Book

func SerializeBook(bookModel models.Book) Book {
	return Book{Name: bookModel.Name, Author: bookModel.Author, Publication: bookModel.Publication}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBools()
	serializedBooks := make([]Book, len(newBooks))
	for index, modelBook := range newBooks {
		serializedBooks[index] = SerializeBook(modelBook)
	}
	resp, _ := json.Marshal(serializedBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBoolById(id)

	serilizedBook := SerializeBook(bookDetails)

	resp, _ := json.Marshal(serilizedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	serilizedBook := SerializeBook(b)
	resp, _ := json.Marshal(serilizedBook)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(id)
	serilizedBook := SerializeBook(book)
	resp, _ := json.Marshal(serilizedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBoolById(id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	serilizedBook := SerializeBook(bookDetails)

	resp, _ := json.Marshal(serilizedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
