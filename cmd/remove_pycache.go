package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tangjicheng46/op/utils"
	"os"
)

var removePycacheCmd = &cobra.Command{
	Use:   "remove_pycache",
	Short: "a",
	Long:  "b",
	Run: func(cmd *cobra.Command, args []string) {
		errString := "usage: op remove_pycache input_directory"
		if len(args) != 1 {
			fmt.Println(errString)
			os.Exit(1)
		}
		if err := utils.RemoveSpecific(args[0], "__pycache__"); err != nil {
			fmt.Println(err)
			fmt.Println(errString)
			os.Exit(1)
		}
	},
}
