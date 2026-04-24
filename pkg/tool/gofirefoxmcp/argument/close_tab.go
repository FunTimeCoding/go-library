package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type CloseTab struct {
	TabIdentifier request.Integer `json:"tab_id"`
}
