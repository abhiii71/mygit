package cmd

import (
	"fmt"
	"mygit/internal/core"

	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch [name]",
	Short: "Manage branches",
	Long:  `Create, list, or delete branches in the repository.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			core.ListBranches()
		} else {
			branchName := args[0]
			err := core.CreateBranch(branchName)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			fmt.Printf("Branch '%s' created successfully.\n", branchName)
		}
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
