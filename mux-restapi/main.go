package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// creating a store at package level
var store Store = Store{}

type Store struct {
	Books []Book
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&store.Books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range store.Books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(&item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(1000000))
	store.Books = append(store.Books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.Books {
		if item.Id == params["id"] {
			store.Books = append(store.Books[:index], store.Books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = strconv.Itoa(rand.Intn(1000000))
			store.Books = append(store.Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&store.Books)

}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.Books {
		if item.Id == params["id"] {
			store.Books = append(store.Books[:index], store.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(&store.Books)
}

func main() {
	fmt.Println("hello rest GO")

	//create some data
	store.Books = append(store.Books, Book{Id: "1", Name: "book1"})
	store.Books = append(store.Books, Book{Id: "2", Name: "book2"})
	store.Books = append(store.Books, Book{Id: "3", Name: "book3"})
	store.Books = append(store.Books, Book{Id: "4", Name: "book4"})
	store.Books = append(store.Books, Book{Id: "5", Name: "book5"})

	router := mux.NewRouter()

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8011", router))

}
