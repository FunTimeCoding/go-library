package telemetry

import "net/http"

type Client struct {
	base   string
	client *http.Client
}
