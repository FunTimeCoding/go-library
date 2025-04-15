package issues

import (
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic_client/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func FromResponseScheme(r *models.ResponseScheme) *Issues {
	result := New()
	var response issue.Response
	errors.PanicOnError(json.Unmarshal(r.Bytes.Bytes(), &response))

	for _, v := range response.Values {
		result.Add(issue.FromValue(v))
	}

	return result
}
