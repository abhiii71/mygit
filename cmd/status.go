package cmd

import (
	"fmt"
	"mygit/internal/core"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the status of the repository",
	Long:  `Shows the current status of the repository, including staged, unstaged and untracked files.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := core.Status()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
