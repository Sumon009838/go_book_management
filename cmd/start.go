/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"BookServer/authentication"
	"BookServer/data"
	"BookServer/handler"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var (
	port     string
	username string
	password string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		address := "8000"
		if len(port) > 0 {
			address = port
		}
		if len(username) > 0 {
			fmt.Println(username)
			fmt.Println(password)
			data.Creds = map[string]data.Credential{
				username: {Username: username, Password: password},
			}
		} else {
			data.Creds = map[string]data.Credential{
				"suman": {Username: "suman", Password: "123"},
			}
		}
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
		fmt.Println("Listening and Serving to ", address)
		err := http.ListenAndServe("localhost:"+address, r)
		if err != nil {
			log.Fatalln(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&port, "port", "n", "", "Server port")
	startCmd.Flags().StringVarP(&username, "username", "u", "", "Username to login")
	startCmd.Flags().StringVarP(&password, "password", "p", "", "Password to login")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
