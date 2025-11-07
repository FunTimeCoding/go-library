package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/customer"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/integers"
)

func (c *Client) CreateCustomerIssue(
	desk int,
	requestType int,
	summary string,
	description string,
) *models.CustomerRequestScheme {
	result, _, e := c.service.Request.Create(
		c.context,
		&models.CreateCustomerRequestPayloadScheme{
			ServiceDeskID: integers.ToString(desk),
			RequestTypeID: integers.ToString(requestType),
			RequestFieldValues: map[string]interface{}{
				customer.SummaryField:     summary,
				customer.DescriptionField: description,
			},
		},
	)
	errors.PanicOnError(e)

	return result
}
