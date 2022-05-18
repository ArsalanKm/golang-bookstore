package main

import (
	"ArsalanKm/golang-bookstore/package/routes"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
