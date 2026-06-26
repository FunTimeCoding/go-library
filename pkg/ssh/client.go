package ssh

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	client *ssh.Client
	sftp   *sftp.Client
	Panic  bool
}
