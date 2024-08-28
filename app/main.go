package main

import (
	"fmt"

	"github.com/mauricioww/goauth/app/utils"
)

func main() {
	password := "secret"
	hash, _ := utils.HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match, _ := utils.ValidatePassword(password, hash)
	fmt.Println("Match:   ", match)
}
