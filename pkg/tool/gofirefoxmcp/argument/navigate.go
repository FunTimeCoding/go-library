package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type Navigate struct {
	TabIdentifier request.Integer `json:"tab_id"`
	Locator       string          `json:"url"`
}
