package grafana

import "github.com/grafana/grafana-openapi-client-go/client"

type Client struct {
	client *client.GrafanaHTTPAPI
}
