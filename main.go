package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Genre  string  `json:"genre"`
	Author *Author `json:"author"`
}
type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var Books map[string]Book
var Authors map[string]Author

func getbooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	for _, book := range Books {
		books = append(books, book)
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func getbook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, ok := Books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func addbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := Books[book.ID]
	if ok {
		http.Error(w, "Book already exists", http.StatusConflict)
		return
	}
	Books[book.ID] = book
	Authors[book.Author.ID] = *book.Author
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func updatebook(w http.ResponseWriter, r *http.Request) {
	oldid := chi.URLParam(r, "id")
	if len(oldid) == 0 {
		http.Error(w, "No ID provided", http.StatusBadRequest)
		return
	}
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := Books[oldid]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if book.ID != oldid {
		http.Error(w, "You cannot change ID of a book.", http.StatusBadRequest)
		return
	}
	Books[oldid] = book
	Authors[book.Author.ID] = *book.Author
	_, err = w.Write([]byte("Book updated successfully"))
	if err != nil {
		http.Error(w, "Can not Write Data", http.StatusInternalServerError)
		return
	}
}
func deletebook(w http.ResponseWriter, r *http.Request) {
	oldid := chi.URLParam(r, "id")
	if len(oldid) == 0 {
		http.Error(w, "No ID provided", http.StatusBadRequest)
		return
	}
	_, ok := Books[oldid]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	delete(Authors, Books[oldid].Author.ID)
	delete(Books, oldid)
	_, err := w.Write([]byte("Data Deleted successfully"))
	if err != nil {
		http.Error(w, "Can not Write Data", http.StatusInternalServerError)
		return
	}
}
func getauthors(w http.ResponseWriter, r *http.Request) {
	var authors []Author
	for _, author := range Authors {
		authors = append(authors, author)
	}
	err := json.NewEncoder(w).Encode(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func getauthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author, ok := Authors[id]
	if !ok {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func find_by_genre(w http.ResponseWriter, r *http.Request) {
	gen := chi.URLParam(r, "genre")
	if len(gen) == 0 {
		http.Error(w, "No genre provided", http.StatusBadRequest)
	}
	for _, book := range Books {
		if book.Genre == gen {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
func init1() {
	Books = map[string]Book{
		"1": {ID: "1", Title: "Movie1", Genre: "Comedy", Author: &Author{ID: "1", FirstName: "Suman", LastName: "Sarker"}},
		"2": {ID: "2", Title: "Movie2", Genre: "Comedy", Author: &Author{ID: "2", FirstName: "Hamim", LastName: "Hossain"}},
	}
	Authors = map[string]Author{
		"1": {FirstName: "Suman", LastName: "Sarker"},
		"2": {FirstName: "Hamim", LastName: "Hossain"},
	}
}
func main() {
	init1()
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Get("/", getbooks)
			r.Get("/{id}", getbook)
			r.Post("/", addbook)
			r.Put("/{id}", updatebook)
			r.Delete("/{id}", deletebook)
		})
		r.Route("/authors", func(r chi.Router) {
			r.Get("/", getauthors)
			r.Get("/{id}", getauthor)
		})
		r.Get("/find/{genre}", find_by_genre)
	})
	fmt.Println("Listening and Serving to 8000")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
