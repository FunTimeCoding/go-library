package request_context

import "github.com/funtimecoding/go-library/pkg/web/constant"

func (c *Context) Scheme() string {
	if p := c.request.Header.Get(constant.ForwardedProtocol); p != "" {
		return p
	}

	if c.request.TLS != nil {
		return constant.Secure
	}

	return constant.Insecure
}
