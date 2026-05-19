package conversations

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
)

type Server struct {
	claude *claude.Client
	store  *store.Store
}
