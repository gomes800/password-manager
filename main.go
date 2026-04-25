package main

import (
	"fmt"

	"github.com/gomes800/password-manager/security"
)

func main() {
	password := "aleatoriosenha12345"

	hash, salt, err := security.HashPassword(password)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hash: %x\n", hash)
	fmt.Printf("Salt: %x\n", salt)
}
