package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tangjicheng46/op/utils"
)

var removePycacheCmd = &cobra.Command{
	Use:   "remove_pycache",
	Short: "a",
	Long:  "b",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		// RemovePycache(args[1])
	},
}

func removePycache(root string) error {
	return utils.RemoveSpecific(root, "__pycache__")
}
