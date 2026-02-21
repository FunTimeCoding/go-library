package mock_client

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

type Client struct {
	alerts []*alert.Alert
}
