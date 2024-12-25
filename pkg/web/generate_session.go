package web

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func GenerateSession() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	errors.PanicOnError(err)

	return hex.EncodeToString(b)
}
