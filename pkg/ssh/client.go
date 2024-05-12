package ssh

import "golang.org/x/crypto/ssh"

type Client struct {
	client *ssh.Client
}
