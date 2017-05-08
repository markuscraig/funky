package main

import (
	"fmt"

	"github.com/tj/go-prompt"
)

func main() {
	println("need your name!")
	first := prompt.String("First name: ")
	last := prompt.String("Last name: ")
	fmt.Printf("\nHello %s %s\n", first, last)
}
