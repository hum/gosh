package main

import (
	"fmt"
	"os"

	"github.com/hum/gosh/shell"
)

func main() {
	if err := shell.Execute(); err != nil {
		fmt.Printf("gosh: %s\n", err)
		os.Exit(0)
	}
}
