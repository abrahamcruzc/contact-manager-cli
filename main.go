/*
Copyright © 2025 ABRAHAM CRUZ abrahamcz_@hotmail.com
*/
package main

import (
	"fmt"
	"log"

	"github.com/abrahamcruzc/contact-manager-cli/cmd"
	"github.com/abrahamcruzc/contact-manager-cli/database"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database")
	}
	
	fmt.Println("📔 Contact management system initialized.")
	cmd.Execute()
}
