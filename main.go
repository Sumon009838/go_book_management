package main

import (
	"BookServer/authentication"
	"BookServer/data"
	"BookServer/handler"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
)

func main() {
	data.Init1()
	r := chi.NewRouter()
	r.Get("/ping", handler.Pong)
	r.Post("/login", authentication.Login)
	r.Get("/logout", authentication.Logout)
	r.Group(func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Get("/", handler.Getbooks)
			r.Get("/{id}", handler.Getbook)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(data.TokenAuth))
				r.Use(jwtauth.Authenticator(data.TokenAuth))
				r.Post("/", handler.Addbook)
				r.Put("/{id}", handler.Updatebook)
				r.Delete("/{id}", handler.Deletebook)
			})
		})
		r.Route("/authors", func(r chi.Router) {
			r.Get("/", handler.Getauthors)
			r.Get("/{id}", handler.Getauthor)
		})
		r.Route("/find", func(r chi.Router) {
			r.Get("/{genre}", handler.Find_by_genre)
		})
	})
	fmt.Println("Listening and Serving to 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
