package main

import (
	"base-setup/internal/wiring"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "app",
	Short: "Simple app",
	Long:  "This command is used for starting simple app",
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "start as api server.",
	Long:  `start as api server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		godotenv.Load(".env")

		// set time zone
		loc, _ := time.LoadLocation("Asia/Tokyo")
		time.Local = loc

		app, cleanup, err := wiring.InitializeStandaloneServer()
		if err != nil {
			log.Fatal(err)
		}

		defer cleanup()

		return app.Start()
	},
}

func main() {
	rootCmd.AddCommand(startCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
