package cmd

import (
	"fmt"
	"mygit/internal/core"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [files]",
	Short: "Add file content to the index",
	Long:  "Add the specified files to the index (tracked files).",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Add(args); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Files added to the index.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
