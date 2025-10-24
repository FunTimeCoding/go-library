package chromium

import (
	"github.com/funtimecoding/go-library/pkg/chromium/tab"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) Tabs() []*tab.Tab {
	var result []*tab.Tab
	notation.DecodeStrict(
		web.GetString(
			web.InsecureClient(),
			locator.New(
				c.host,
			).Port(c.port).Path("/json").Insecure().String(),
		),
		&result,
		true,
	)

	return result
}
