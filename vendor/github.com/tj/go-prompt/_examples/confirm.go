package main

import "github.com/tj/go-prompt"

func main() {
	if ok := prompt.Confirm("launch %s? ", "something"); ok {
		println("launching")
	} else {
		println("not launching")
	}
}
