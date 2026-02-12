package proxmox

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/proxmox/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/funtimecoding/go-library/pkg/web/verbose_transport"
	"github.com/luthermonson/go-proxmox"
	"net/http"
)

func New(
	host string,
	o ...Option,
) *Client {
	// https://pve.proxmox.com/pve-docs/api-viewer
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	var option []proxmox.Option

	if result.user != "" && result.password != "" {
		option = append(
			option,
			proxmox.WithCredentials(
				&proxmox.Credentials{
					Username: result.user,
					Password: result.password,
				},
			),
		)
	} else if result.token != "" && result.secret != "" {
		option = append(
			option,
			proxmox.WithAPIToken(result.token, result.secret),
		)
	} else {
		panic("missing credentials")
	}

	if result.log {
		option = append(
			option,
			proxmox.WithLogger(
				&proxmox.LeveledLogger{Level: proxmox.LevelDebug},
			),
		)
	}

	var c *http.Client

	if result.verbose {
		c = verbose_transport.New(!result.selfSigned)
	} else if result.selfSigned {
		c = web.InsecureClient()
	}

	if c != nil {
		option = append(option, proxmox.WithHTTPClient(c))
	}

	result.client = proxmox.NewClient(
		locator.New(host).Port(constant.Port).Base(constant.Base).String(),
		option...,
	)

	return result
}
