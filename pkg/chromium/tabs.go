package chromium

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium/tab"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func (c *Client) Tabs() []*tab.Tab {
	var result []*tab.Tab
	notation.DecodeStrict(
		web.GetString(
			web.Client(false),
			fmt.Sprintf(
				"%s://%s:%d/json",
				constant.InsecureScheme,
				c.host,
				c.port,
			),
		),
		&result,
		true,
	)

	return result
}
