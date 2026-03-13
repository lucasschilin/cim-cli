package main

import (
	"fmt"

	"github.com/lucasschilin/commit-improver-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
