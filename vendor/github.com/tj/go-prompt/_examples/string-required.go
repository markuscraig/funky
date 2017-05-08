package main

import (
	"fmt"

	"github.com/tj/go-prompt"
)

func main() {
	first := prompt.StringRequired("First name: ")
	last := prompt.StringRequired("Last name: ")
	fmt.Printf("\nHello %q %q\n", first, last)
}
