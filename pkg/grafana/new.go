package grafana

import (
	"crypto/tls"
	"github.com/go-openapi/strfmt"
	"github.com/grafana/grafana-openapi-client-go/client"
	"net/url"
	"os"
)

func New() *Client {
	// TODO: Stub, implement me
	return &Client{
		client: client.NewHTTPClientWithConfig(
			strfmt.Default,
			&client.TransportConfig{
				// Host is the domain name or IP address of the host that serves the API.
				Host: "localhost:3000",
				// BasePath is the URL prefix for all API paths, relative to the host root.
				BasePath: "/api",
				// Schemes are the transfer protocols used by the API (http or https).
				Schemes: []string{"http"},
				// APIKey is an optional API key or service account token.
				APIKey: os.Getenv("API_ACCESS_TOKEN"),
				// BasicAuth is optional basic auth credentials.
				BasicAuth: url.UserPassword("admin", "admin"),
				// OrgID provides an optional organization ID.
				// OrgID is only supported with BasicAuth since API keys are already org-scoped.
				OrgID: 1,
				// TLSConfig provides an optional configuration for a TLS client
				TLSConfig: &tls.Config{},
				// NumRetries contains the optional number of attempted retries
				NumRetries: 3,
				// RetryTimeout sets an optional time to wait before retrying a request
				RetryTimeout: 0,
				// RetryStatusCodes contains the optional list of status codes to retry
				// Use "x" as a wildcard for a single digit (default: [429, 5xx])
				RetryStatusCodes: []string{"420", "5xx"},
				// HTTPHeaders contains an optional map of HTTP headers to add to each request
				HTTPHeaders: map[string]string{},
			},
		),
	}
}
