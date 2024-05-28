package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func KnownHosts() ssh.HostKeyCallback {
	result, e := knownhosts.New(
		system.Join(system.Home(), ConfigurationDirectory, KnownHostsFile),
	)
	errors.PanicOnError(e)

	return result
}
