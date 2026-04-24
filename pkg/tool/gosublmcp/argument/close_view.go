package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type CloseView struct {
	ViewIdentifier request.Integer `json:"view_id"`
}
