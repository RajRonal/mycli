package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"mcli/database"
)

var Up *cobra.Command

func init() {
	Up = &cobra.Command{
		Use:   "up",
		Short: "migration is up",
		Long:  " ",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")
			err := database.ConnectAndMigrate("localhost", "5432", "audiophile", "local", "local", database.SSLModeDisable)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.AddCommand(Up)
}
