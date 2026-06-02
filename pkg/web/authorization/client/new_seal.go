package client

import "crypto/cipher"

func newSeal(block cipher.Block) (cipher.AEAD, error) {
	return cipher.NewGCM(block)
}
