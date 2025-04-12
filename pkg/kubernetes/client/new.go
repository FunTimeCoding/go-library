package client

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/client_configuration"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/client_context"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/metrics"
)

func New(clusters []string) *Client {
	result := Stub()
	result.context = context.Background()
	result.configuration = client_configuration.New()
	result.client = client.New(result.configuration)
	result.metric = metrics.New(result.configuration)
	result.cluster = client_context.Current()

	if len(clusters) == 0 {
		clusters = result.ConfigurationContexts(true)
	}

	for _, element := range clusters {
		c, e := NewContext(element)

		if e != nil {
			continue
		}

		result.clients[element] = c
	}

	return result
}
