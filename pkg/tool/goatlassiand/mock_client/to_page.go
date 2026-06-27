package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func toPage(r *response.Page) *page.Page {
	return page.New(r, "https://mock.atlassian.net")
}
