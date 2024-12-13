package cmd

import (
	"fmt"
	"mygit/internal/core"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [files...]",
	Short: "Add files to the staging area",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := core.AddToTrackedFiles(args)
		if err != nil {
			fmt.Println("Error adding files:", err)
			os.Exit(1)
		}
		fmt.Println("Files added to the staging area.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
