package response

import "github.com/funtimecoding/go-library/pkg/iterm/session"

type Sessions struct {
	Sessions []session.Session `json:"sessions"`
}
