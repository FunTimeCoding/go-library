package telemetry

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func New() *Client {
	host := os.Getenv(HostEnvironment)
	port := os.Getenv(PortEnvironment)

	if host == "" || port == "" {
		return nil
	}

	return &Client{
		base: fmt.Sprintf("http://%s:%s", host, port),
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
