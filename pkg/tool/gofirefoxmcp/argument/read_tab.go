package argument

import (
	"github.com/funtimecoding/go-library/pkg/generative/mark/request/boolean"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request/integer"
)

type ReadTab struct {
	TabIdentifier integer.Integer `json:"tab_id"`
	Raw           boolean.Boolean `json:"raw"`
}
