package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request/integer"

type SendKey struct {
	SessionIdentifier string          `json:"session_id"`
	Keys              []string        `json:"keys"`
	Interval          integer.Integer `json:"interval"`
}
