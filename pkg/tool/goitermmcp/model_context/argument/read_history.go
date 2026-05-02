package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request/integer"

type ReadHistory struct {
	SessionIdentifier string          `json:"session_id"`
	Count             integer.Integer `json:"count"`
}
