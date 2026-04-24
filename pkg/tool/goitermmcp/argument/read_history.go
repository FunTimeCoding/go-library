package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type ReadHistory struct {
	SessionIdentifier string          `json:"session_id"`
	Count             request.Integer `json:"count"`
}
