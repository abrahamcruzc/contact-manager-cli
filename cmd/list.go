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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contacts",
	Run: func(cmd *cobra.Command, args []string) {
		var contacts []models.Contact
		database.DB.Find(&contacts)
		
		for _, contact := range contacts {
			fmt.Printf("ID: %v \t NAME: %v \t EMAIL: %v \t PHONE: %v \n", contact.ID, contact.Name, contact.Email, contact.Phone)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
