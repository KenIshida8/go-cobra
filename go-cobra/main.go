package main

import (
	"fmt"
	"os"

	"go-CLItool-test/go-cobra/cmd"
	)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

