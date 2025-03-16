package grafana

import goapi "github.com/grafana/grafana-openapi-client-go/client"

type Client struct {
	client *goapi.GrafanaHTTPAPI
}
