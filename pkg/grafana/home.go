package grafana

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func (c *Client) Home() *models.GetHomeDashboardResponse {
	r, e := c.client.Dashboards.GetHomeDashboard()
	errors.PanicOnError(e)

	return r.Payload
}
