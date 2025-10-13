package customer

import "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"

func NewSlice(v []*models.CustomerRequestScheme) []*Issue {
	var result []*Issue

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
