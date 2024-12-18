package security

import (
	"fmt"

	"github.com/xyproto/randomstring"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(rawPassword string, salt string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(salt+rawPassword+salt), 14)
	if err != nil {
		return nil, fmt.Errorf("generate password hash: %w", err)
	}
	hashString := string(hash)
	return &hashString, nil
}

func GenSalt() string {
	return randomstring.String(16)
}
