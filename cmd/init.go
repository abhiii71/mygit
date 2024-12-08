package cmd

import (
	"fmt"
	"mygit/internal/core"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Long:  "Creates a new MyGit repository in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		err := core.InitializeRepo()
		if err != nil {
			fmt.Printf("Error initializing repository: %v\n", err)
		} else {
			fmt.Println("Initialized an empty MyGit repository.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
