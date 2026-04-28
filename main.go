package main

import (
	"fmt"

	"github.com/gomes800/password-manager/security"
)

func main() {
	password1 := "aleatoriosenha123456"
	bankPassword := []byte("senhasegura12345@")

	password2 := "churinchurin"
	lifePassword := []byte("senhahipersegura123456#")

	key1, salt1, err := security.HashPassword(password1)
	if err != nil {
		panic(err)
	}

	key2, salt2, err := security.HashPassword(password2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("key1: %x\n", key1)
	fmt.Printf("salt1: %x\n", salt1)

	fmt.Printf("key2: %x\n", key2)
	fmt.Printf("salt2: %x\n", salt2)

	ciphertext, nonce, err := security.Encrypt(key1, bankPassword)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %x\n", ciphertext)

	plaintext, err := security.Decrypt(key1, ciphertext, nonce)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext)

	ciphertext2, nonce2, err := security.Encrypt(key2, lifePassword)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %x\n", ciphertext2)

	plaintext2, err := security.Decrypt(key1, ciphertext2, nonce2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext2)
}
