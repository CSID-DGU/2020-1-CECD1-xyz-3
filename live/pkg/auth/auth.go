package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

const saltSize = 32

func NewToken(email string) (string, error) {
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s-%s", email, string(salt))))
	data := hash.Sum(nil)
	return hex.EncodeToString(data), nil
}
