package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request/integer"

type NavigateBack struct {
	TabIdentifier integer.Integer `json:"tab_id"`
}
