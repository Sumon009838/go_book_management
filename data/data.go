package data

import (
	"github.com/go-chi/jwtauth/v5"
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
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var mySigningKey = []byte("secret")
var Books map[string]Book
var Authors map[string]Author
var Creds map[string]Credential
var TokenAuth *jwtauth.JWTAuth

func Init1() {
	TokenAuth = jwtauth.New("HS256", mySigningKey, nil)
	Books = map[string]Book{
		"1": {ID: "1", Title: "Movie1", Genre: "comedy", Author: &Author{ID: "1", FirstName: "Suman", LastName: "Sarker"}},
		"2": {ID: "2", Title: "Movie2", Genre: "comedy", Author: &Author{ID: "2", FirstName: "Hamim", LastName: "Hossain"}},
	}
	Authors = map[string]Author{
		"1": {FirstName: "Suman", LastName: "Sarker"},
		"2": {FirstName: "Hamim", LastName: "Hossain"},
	}
	Creds = map[string]Credential{
		"suman": {Username: "suman", Password: "pass1"},
		"hamim": {Username: "hamim", Password: "pass2"},
	}
}
