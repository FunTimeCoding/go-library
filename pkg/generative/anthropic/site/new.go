package site

import "github.com/funtimecoding/go-library/pkg/chromium/protocol"

func New() *Site {
	return &Site{protocol: protocol.New("claude.ai/settings/usage")}
}
