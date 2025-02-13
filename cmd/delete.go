/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/abrahamcruzc/contact-manager-cli/database"
	"github.com/abrahamcruzc/contact-manager-cli/models"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete contact by id",
	Run: func(cmd *cobra.Command, args []string) {
		if id == 0 {
			cmd.Println("Error: invalid ID")
			return
		}
		
		result := database.DB.Delete(&models.Contact{}, id)
		
		if result.RowsAffected == 0 {
			cmd.Println("Error: No contact was found with that ID")
		} else {
			cmd.Println("Contact deleted successfully")
		}
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "Contact ID")
	rootCmd.AddCommand(deleteCmd)
}
