package conversations

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
)

func New(
	c *claude.Client,
	s *store.Store,
) *Server {
	return &Server{claude: c, store: s}
}
