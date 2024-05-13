package secure_shell

import (
	"bytes"
	"golang.org/x/crypto/ssh"
)

func SessionBuffers(s *ssh.Session) (*bytes.Buffer, *bytes.Buffer) {
	var outputBuffer bytes.Buffer
	var errorBuffer bytes.Buffer
	s.Stdout = &outputBuffer
	s.Stderr = &errorBuffer

	return &outputBuffer, &errorBuffer
}
