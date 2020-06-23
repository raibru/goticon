package main

import (
	"os"

	"github.com/raibru/pktfmt/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
