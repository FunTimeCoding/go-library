package ssh

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/secure_shell"
	"golang.org/x/crypto/ssh"
	"net"
)

func NewWithFile(
	user string,
	host string,
	keyPath string,
	secure bool,
) *Client {
	var callback func(
		hostname string,
		remote net.Addr,
		k ssh.PublicKey,
	) error

	if secure {
		callback = secure_shell.KnownHosts()
	} else {
		callback = ssh.InsecureIgnoreHostKey()
	}

	return &Client{
		client: secure_shell.Dial(
			fmt.Sprintf("%s:22", host),
			&ssh.ClientConfig{
				User: user,
				Auth: []ssh.AuthMethod{
					ssh.PublicKeys(
						secure_shell.ParsePrivateKey(
							system.ReadBytes(keyPath),
						),
					),
				},
				HostKeyCallback: callback,
			},
		),
	}
}
