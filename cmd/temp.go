package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tangjicheng46/op/utils"
)

var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Just for temp test",
	Run: func(cmd *cobra.Command, args []string) {
		inferMain(args)
	},
}

func inferMain(args []string) {
	if len(args) != 2 {
		fmt.Println("usage: op temp data.json your_url")
		return
	}
	fmt.Printf("[tangjicheng] %v\n", args)
	utils.Infer(args[0], args[1])
}
