package loki

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/message"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	telemetry "github.com/funtimecoding/go-library/pkg/web/telemetry/constant"
	"strings"
)

func formatContent(
	v *message.Message,
	f *option.Format,
) string {
	route := v.Value(telemetry.Route)
	body := strings.TrimSpace(v.Value(telemetry.Body))

	if route != "" && body != "" {
		return join.Empty(formatRoute(route, f), " ", body)
	}

	if route != "" {
		return formatRoute(route, f)
	}

	if body != "" {
		return body
	}

	if m := v.Value(message.SlogMessage); m != "" {
		return m
	}

	if v.Text != "" {
		return strings.TrimSpace(v.Text)
	}

	return ""
}
