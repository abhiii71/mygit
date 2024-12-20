package cmd

import (
	"fmt"
	"mygit/internal/core"
	"os"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Long:  "Display a list of all commits made in the repository in reverse chronological order.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.ShowLog(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
