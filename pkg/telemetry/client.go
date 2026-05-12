package telemetry

import "net/http"

type Client struct {
	base       string
	httpClient *http.Client
}
