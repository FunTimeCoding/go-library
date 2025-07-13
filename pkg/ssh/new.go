package ssh

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ssh/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/secure_shell"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
)

func New(
	user string,
	host string,
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
						secure_shell.Signers(
							system.UnixSocket(
								os.Getenv(constant.SocketEnvironment),
							),
						)...,
					),
				},
				HostKeyCallback: callback,
			},
		),
	}
}
