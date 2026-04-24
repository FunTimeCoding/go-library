package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type ReadView struct {
	ViewIdentifier request.Integer `json:"view_id"`
}
