package handler

import (
	"BookServer/data"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Getbooks(w http.ResponseWriter, r *http.Request) {
	var books []data.Book
	for _, book := range data.Books {
		books = append(books, book)
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func Getbook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, ok := data.Books[id]
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
func Addbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book data.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := data.Books[book.ID]
	if ok {
		http.Error(w, "Book already exists", http.StatusConflict)
		return
	}
	data.Books[book.ID] = book
	data.Authors[book.Author.ID] = *book.Author
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Updatebook(w http.ResponseWriter, r *http.Request) {
	oldid := chi.URLParam(r, "id")
	if len(oldid) == 0 {
		http.Error(w, "No ID provided", http.StatusBadRequest)
		return
	}
	var book data.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := data.Books[oldid]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if book.ID != oldid {
		http.Error(w, "You cannot change ID of a book.", http.StatusBadRequest)
		return
	}
	data.Books[oldid] = book
	data.Authors[book.Author.ID] = *book.Author
	_, err = w.Write([]byte("Book updated successfully"))
	if err != nil {
		http.Error(w, "Can not Write Data", http.StatusInternalServerError)
		return
	}
}
func Deletebook(w http.ResponseWriter, r *http.Request) {
	oldid := chi.URLParam(r, "id")
	if len(oldid) == 0 {
		http.Error(w, "No ID provided", http.StatusBadRequest)
		return
	}
	_, ok := data.Books[oldid]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	delete(data.Authors, data.Books[oldid].Author.ID)
	delete(data.Books, oldid)
	_, err := w.Write([]byte("Data Deleted successfully"))
	if err != nil {
		http.Error(w, "Can not Write Data", http.StatusInternalServerError)
		return
	}
}
func Getauthors(w http.ResponseWriter, r *http.Request) {
	var authors []data.Author
	for _, author := range data.Authors {
		authors = append(authors, author)
	}
	err := json.NewEncoder(w).Encode(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func Getauthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author, ok := data.Authors[id]
	if !ok {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Find_by_genre(w http.ResponseWriter, r *http.Request) {
	gen := chi.URLParam(r, "genre")
	if len(gen) == 0 {
		http.Error(w, "No genre provided", http.StatusBadRequest)
	}
	for _, book := range data.Books {
		if book.Genre == gen {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
