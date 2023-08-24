package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "op",
	Short: "Streamlining daily tasks for seamless computer productivity",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(removePycacheCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tempCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
