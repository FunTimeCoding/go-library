package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type NavigateBack struct {
	TabIdentifier request.Integer `json:"tab_id"`
}
