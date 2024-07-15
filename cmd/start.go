/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"BookServer/data"
	"BookServer/handler"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	port     string
	username string
	password string
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This command starts the server",
	Long:  `Through this cli command we can specify the port, username and password by using flags`,
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
		handler.Allfunc(address)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&port, "port", "n", "", "Server port")
	startCmd.Flags().StringVarP(&username, "username", "u", "", "Username to login")
	startCmd.Flags().StringVarP(&password, "password", "p", "", "Password to login")
}
