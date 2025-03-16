package loki

import "github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client"

type Client struct {
	basic *basic_client.Client
}
