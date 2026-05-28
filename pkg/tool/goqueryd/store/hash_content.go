package store

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashContent(content string) string {
	h := sha256.Sum256([]byte(content))

	return hex.EncodeToString(h[:])
}
