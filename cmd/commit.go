package cmd

import (
	"fmt"
	"mygit/internal/core"
	"os"

	"github.com/spf13/cobra"
)

var commitMessage string

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record Changes to the repository",
	Long:  "Save the current state of tracked files in the repository, with a commit message.",
	Run: func(cmd *cobra.Command, args []string) {
		if commitMessage == "" {
			fmt.Println("Error: Commit message is required. Use the -m flag.")
			os.Exit(1)
		}

		if err := core.Commit(commitMessage); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Commit created successfully.")
	},
}

func init() {
	commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "commit message")
	rootCmd.AddCommand(commitCmd)
}
