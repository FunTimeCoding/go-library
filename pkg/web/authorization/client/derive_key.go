package client

import "crypto/sha256"

func DeriveKey(secret string) []byte {
	h := sha256.Sum256([]byte(secret))

	return h[:]
}
