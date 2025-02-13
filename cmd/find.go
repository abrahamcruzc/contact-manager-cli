/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/abrahamcruzc/contact-manager-cli/database"
	"github.com/abrahamcruzc/contact-manager-cli/models"
	"github.com/spf13/cobra"
)

var id int

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a contact by ID",

	Run: func(cmd *cobra.Command, args []string) {
		if id == 0 {
			cmd.Println("Error: invalid ID")
			return
		}
		var contact models.Contact
		result := database.DB.First(&contact, id)
		
		if result.RowsAffected == 0 {
			cmd.Println("Error: contact was not foud with that ID")
		}
		
		cmd.Printf("ID: %v \t NAME: %v \t EMAIL: %v \t PHONE: %v \n", contact.ID, contact.Name, contact.Email, contact.Phone)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().IntVarP(&id, "id", "i", 0, "Find by id")
}
