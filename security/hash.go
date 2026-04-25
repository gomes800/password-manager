package security

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) ([]byte, []byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, nil, err
	}

	var (
		time    uint32 = 1
		memory  uint32 = 64 * 1024
		threads uint8  = 4
		keyLen  uint32 = 32
	)

	key := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	return key, salt, nil
}
