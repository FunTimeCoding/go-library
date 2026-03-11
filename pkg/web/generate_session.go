package web

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func GenerateSession() string {
	b := make([]byte, 32)
	_, e := rand.Read(b)
	errors.PanicOnError(e)

	return hex.EncodeToString(b)
}
