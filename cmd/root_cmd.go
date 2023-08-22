package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "op",
	Short: "A tool to remove Python cache files",
	Long:  "A simple tool to remove .pyc and __pycache__ directories.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing if no subcommand is specified
		if err := cmd.Help(); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(removePycacheCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
