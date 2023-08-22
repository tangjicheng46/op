package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "op",
	Short: "A tool to remove Python cache files",
	Long:  "A simple tool to remove .pyc and __pycache__ directories.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing if no subcommand is specified
		
		fmt.Printf("[in rootCmd] %v\n", args)

		if err := cmd.Help(); err != nil {
			return
		}
	},
}

func Execute() error {
	rootCmd.AddCommand(removePycacheCmd)
	return rootCmd.Execute()
}
