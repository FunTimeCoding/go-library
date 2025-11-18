package request_context

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web/telemetry/constant"
	"log/slog"
	"net/http"
	"strings"
)

func (c *Context) Attribute() []slog.Attr {
	b := c.Body()
	result := []slog.Attr{
		slog.String(constant.RequestMethod, c.request.Method),
		slog.String(constant.Path, c.request.URL.Path),
		slog.String(constant.Scheme, c.Scheme()),
		slog.String(constant.Query, c.request.URL.RawQuery),
		slog.Int(constant.Status, http.StatusOK),
		slog.String(constant.Client, c.ClientAddress()),
		slog.String(constant.Peer, c.request.RemoteAddr),
		slog.String(constant.Protocol, c.ProtocolVersion()),
		slog.String(constant.Server, c.request.Host),
		slog.String(constant.UserAgent, c.request.UserAgent()),
		slog.Int(constant.BodySize, len(b)),
		slog.String(constant.Body, b),
	}

	for k, v := range c.request.Header {
		result = append(
			result,
			slog.Any(
				key_value.Dot(constant.HeaderPrefix, strings.ToLower(k)),
				v,
			),
		)
	}

	return result
}
