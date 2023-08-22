package main

import (
	"fmt"
	"github.com/tangjicheng46/op/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
