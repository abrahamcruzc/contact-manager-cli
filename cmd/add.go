/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/abrahamcruzc/contact-manager-cli/database"
	"github.com/abrahamcruzc/contact-manager-cli/models"
	"github.com/spf13/cobra"
)

var name, email, phone string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact",
	Run: func(cmd *cobra.Command, args []string) {
		contact := models.Contact{
			Name: name,
			Email: email,
			Phone: phone,
		}
		database.DB.Create(&contact)
		fmt.Println("Contact added: ", &contact)
	},
}

func init() {
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Contact name")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Contact email")
	addCmd.Flags().StringVarP(&phone, "phone", "p", "", "Contact phone")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
	addCmd.MarkFlagRequired("phone")
	
	rootCmd.AddCommand(addCmd)
}
