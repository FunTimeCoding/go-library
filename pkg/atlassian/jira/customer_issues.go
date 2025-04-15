package jira

import (
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	basicIssues "github.com/funtimecoding/go-library/pkg/atlassian/jira/basic_client/issue/issues"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/customer"
	"github.com/funtimecoding/go-library/pkg/errors"
	"slices"
)

func (c *Client) CustomerIssues(all bool) []*customer.Issue {
	issues, r, e := c.service.Request.Gets(
		c.context,
		&models.ServiceRequestOptionScheme{},
		0,
		10,
	)
	errors.PanicOnError(e)
	basicList := basicIssues.FromResponseScheme(r)
	var result []*customer.Issue

	for _, i := range customer.NewSlice(issues.Values) {
		basic := basicList.Key(i.Key)
		i.Title = basic.Title
		i.Body = basic.Body

		if all {
			result = append(result, i)
		} else if !slices.Contains(constant.ServiceDeskDone, i.Status) {
			result = append(result, i)
		}
	}

	return result
}
