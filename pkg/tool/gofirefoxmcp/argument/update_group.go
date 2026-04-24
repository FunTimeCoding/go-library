package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type UpdateGroup struct {
	GroupIdentifier request.Integer `json:"group_id"`
	Title           string          `json:"title"`
	Color           string          `json:"color"`
}
