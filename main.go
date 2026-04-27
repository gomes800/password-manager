package main

import (
	"fmt"

	"github.com/gomes800/password-manager/security"
)

func main() {
	password := "aleatoriosenha123456"
	bankPassword := []byte("senhasegura12345@")

	key, salt, err := security.HashPassword(password)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Key: %x\n", key)
	fmt.Printf("Salt: %x\n", salt)

	ciphertext, nonce, err := security.Encrypt(key, bankPassword)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %x\n", ciphertext)

	plaintext, err := security.Decrypt(key, ciphertext, nonce)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext)
}
