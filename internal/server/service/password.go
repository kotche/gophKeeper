package service

import (
	"crypto/sha1"
	"fmt"
)

// PasswordEncryptor encrypts passwords
type PasswordEncryptor struct {
	keyPassword string
}

func NewPasswordEncryptor(keyPassword string) *PasswordEncryptor {
	return &PasswordEncryptor{
		keyPassword: keyPassword,
	}
}

// GeneratePasswordHash generates an encrypted password
func (p *PasswordEncryptor) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(p.keyPassword)))
}
