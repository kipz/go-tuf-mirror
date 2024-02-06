package main

import (
	"fmt"
	"os"

	"github.com/docker/go-tuf-mirror/cmd"
)

var version = ""

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
