package main

import "github.com/tj/go-prompt"

func main() {
	{
		name := "Tobi"
		pass := prompt.Password("hi %s enter your password: ", name)
		println(pass)
	}

	{
		pass := prompt.PasswordMasked("masked passwords are cool: ")
		println(pass)
	}
}
