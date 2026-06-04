package client

import (
	"crypto/cipher"
	"github.com/coreos/go-oidc/v3/oidc"
	"sync"
)

type Client struct {
	issuer          string
	identifier      string
	secret          string
	signInPath      string
	callbackLocator string
	seal            cipher.AEAD
	provider        *oidc.Provider
	verifier        *oidc.IDTokenVerifier
	providerOnce    sync.Once
}
