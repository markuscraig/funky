package main

import (
	"github.com/markuscraig/funky/cmd/funky/cmd"
)

var (
	// Version is set during the go build process
	Version string
)

func main() {
	// execute the root command handler
	cmd.Execute(Version)
}
