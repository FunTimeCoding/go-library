package chromium

import "github.com/funtimecoding/go-library/pkg/web"

func NewCombined(combined string) *Client {
	host, port := web.HostPortFromLocatorSplit(combined)

	return New(host, port)
}
