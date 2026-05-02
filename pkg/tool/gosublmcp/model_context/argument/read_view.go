package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request/integer"

type ReadView struct {
	ViewIdentifier integer.Integer `json:"view_id"`
}
