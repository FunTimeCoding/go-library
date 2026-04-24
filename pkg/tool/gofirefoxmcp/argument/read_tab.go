package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type ReadTab struct {
	TabIdentifier request.Integer `json:"tab_id"`
	Raw           request.Boolean `json:"raw"`
}
