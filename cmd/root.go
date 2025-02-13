/*
Copyright © 2025 ABRAHAM CRUZ abrahamcz_@hotmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "contact-manager",
	Short: "Command line contact manager.",
	Long: "CLI application to manage contacts using Go, SQLite and Cobra",

	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("Contact Manager ClI - Use 'contact-manager help' for more information")
	},
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
