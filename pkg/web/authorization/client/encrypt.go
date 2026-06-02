package client

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func (c *Client) encrypt(payload []byte) (string, error) {
	nonce := make([]byte, c.seal.NonceSize())
	_, e := io.ReadFull(rand.Reader, nonce)

	if e != nil {
		return "", e
	}

	sealed := c.seal.Seal(nonce, nonce, payload, nil)

	return base64.RawURLEncoding.EncodeToString(sealed), nil
}
