package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request/integer"

type CloseView struct {
	ViewIdentifier integer.Integer `json:"view_id"`
}
