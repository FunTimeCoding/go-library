package secure_shell

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"net"
)

func HostKnown(host string) bool {
	callback := KnownHosts()
	address := knownhosts.Normalize(host)

	if _, _, e := net.SplitHostPort(address); e != nil {
		address = net.JoinHostPort(address, "22")
	}

	_, private, e := ed25519.GenerateKey(rand.Reader)

	if e != nil {
		return false
	}

	signer, e := ssh.NewSignerFromKey(private)

	if e != nil {
		return false
	}

	e = callback(address, &net.TCPAddr{}, signer.PublicKey())

	if e == nil {
		return true
	}

	var keyError *knownhosts.KeyError

	if errors.As(e, &keyError) {
		return len(keyError.Want) > 0
	}

	return false
}
