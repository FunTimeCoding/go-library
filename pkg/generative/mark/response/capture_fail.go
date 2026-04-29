package response

import (
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/mcp"
)

func CaptureFail(
	h *sentry.Hub,
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return FailAny(
		map[string]any{
			"error":            message,
			"event_identifier": h.CaptureException(e),
		},
	)
}
