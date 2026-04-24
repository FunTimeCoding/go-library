package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type SendKey struct {
	SessionIdentifier string          `json:"session_id"`
	Keys              []string        `json:"keys"`
	Interval          request.Integer `json:"interval"`
}
