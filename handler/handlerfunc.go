package handler

import (
	"BookServer/authentication"
	"BookServer/data"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
)

func Allfunc(address string) {
	r := chi.NewRouter()
	r.Get("/ping", Pong)
	r.Post("/login", authentication.Login)
	r.Get("/logout", authentication.Logout)
	r.Group(func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Get("/", Getbooks)
			r.Get("/{id}", Getbook)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(data.TokenAuth))
				r.Use(jwtauth.Authenticator(data.TokenAuth))
				r.Post("/", Addbook)
				r.Put("/{id}", Updatebook)
				r.Delete("/{id}", Deletebook)
			})
		})
		r.Route("/authors", func(r chi.Router) {
			r.Get("/", Getauthors)
			r.Get("/{id}", Getauthor)
		})
		r.Route("/find", func(r chi.Router) {
			r.Get("/{genre}", Find_by_genre)
		})
	})
	fmt.Println("Listening and Serving to ", address)
	fmt.Println(":" + address)
	err := http.ListenAndServe(":"+address, r)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
