package claude

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"

type SessionToolCount struct {
	Session *session.Session
	Count   int
}
