package grafana

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/go-openapi/strfmt"
	"github.com/grafana/grafana-openapi-client-go/client"
)

func New(
	host string,
	port string,
	token string,
) *Client {
	return &Client{
		client: client.NewHTTPClientWithConfig(
			strfmt.Default,
			&client.TransportConfig{
				Host:     fmt.Sprintf("%s:%s", host, port),
				BasePath: "/api",
				Schemes:  []string{constant.Secure},
				APIKey:   token,
				//BasicAuth: url.UserPassword("admin", "admin"),
				OrgID: 1,
				//TLSConfig: &tls.Config{},
				//NumRetries: 3,
				//RetryTimeout: 0,
				//RetryStatusCodes: []string{"420", "5xx"},
				//HTTPHeaders: map[string]string{},
			},
		),
	}
}
