package tools

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

const passwordChars = "abcdedfghijklmnopqrstABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// SALT ...
// TODO: Move to env variables
const SALT = "MCQc74fe9TKYOJFy3x9E3w2ft/GhkwyjM3O0qRfm7XwdmXuXK+lg718TV2KtUat1"

// RandomPassword ...
func RandomPassword() string {
	b := make([]byte, 8)
	s := ""
	rand.Read(b)
	l := len(passwordChars)
	for _, x := range b {
		s += passwordChars[int(x)%l : int(x)%l+1]
	}
	return s
}

// PasswordHash salts password and turns into sha256 hash
func PasswordHash(password string) string {
	hash := sha256.Sum256([]byte(password + SALT))
	return fmt.Sprintf("%x", hash)
}
