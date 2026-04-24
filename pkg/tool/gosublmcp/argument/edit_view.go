package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type EditView struct {
	ViewIdentifier request.Integer `json:"view_id"`
	OldString      string          `json:"old_string"`
	NewString      string          `json:"new_string"`
	ReplaceAll     request.Boolean `json:"replace_all"`
}
