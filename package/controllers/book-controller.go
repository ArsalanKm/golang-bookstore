package controllers

import (
	"ArsalanKm/golang-bookstore/package/models"
	"ArsalanKm/golang-bookstore/package/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	Book.CreateBook()
	res, _ := json.Marshal(Book)
	w.Header().Set("Content-Type", "application/jon")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/jon")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("invalid id")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	Params := mux.Vars(r)
	ID, err := strconv.ParseInt(Params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("invalid Id on delete book")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	models.DeleteBook(ID)
	res, _ := json.Marshal("deleted")
	w.Write(res)
}
