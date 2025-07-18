package issues

import (
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic_client/issue"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func FromResponseScheme(r *models.ResponseScheme) *Issues {
	var response issue.Response
	notation.DecodeBytesStrict(r.Bytes.Bytes(), &response, true)
	result := New()

	for _, v := range response.Values {
		result.Add(issue.FromValue(v))
	}

	return result
}
