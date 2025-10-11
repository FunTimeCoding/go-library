package jira

import (
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	basic "github.com/funtimecoding/go-library/pkg/atlassian/jira/basic/issue/issues"
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
	list := basic.FromResponseScheme(r)
	var result []*customer.Issue

	for _, i := range customer.NewSlice(issues.Values) {
		b := list.Key(i.Key)

		if b == nil {
			continue
		}

		i.Title = b.Title
		i.Body = b.Body

		if all {
			result = append(result, i)
		} else if !slices.Contains(constant.ServiceDeskDone, i.Status) {
			result = append(result, i)
		}
	}

	return result
}
