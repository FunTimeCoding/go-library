package grafana

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func (c *Client) Dashboards() []*models.PublicDashboardListResponse {
	r, e := c.client.Dashboards.ListPublicDashboards()
	errors.PanicOnError(e)

	return r.Payload.PublicDashboards
}
