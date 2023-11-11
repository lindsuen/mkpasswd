package main

import (
	"fmt"

	s "github.com/lindsuen0/mkpasswd/src"
)

func main() {
	password := new(s.Password)
	s.SetParameter(password)
	s.HandleInputError(password.Length, password.Number, password.Character)

	var counter uint64
	for counter = 0; counter < password.Quantity; counter++ {
		fmt.Println(s.RandomString(password.Length, password.Number, password.Character))
	}
}
