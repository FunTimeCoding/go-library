package client

import (
	"encoding/base64"
	"fmt"
)

func (c *Client) decrypt(encoded string) ([]byte, error) {
	b, e := base64.RawURLEncoding.DecodeString(encoded)

	if e != nil {
		return nil, e
	}

	nonceSize := c.seal.NonceSize()

	if len(b) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	return c.seal.Open(nil, b[:nonceSize], b[nonceSize:], nil)
}
