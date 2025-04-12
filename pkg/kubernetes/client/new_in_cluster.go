package client

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/metrics"
	"k8s.io/client-go/rest"
)

func NewInCLuster(cluster string) *Client {
	configuration, e := rest.InClusterConfig()
	errors.PanicOnError(e)
	result := Stub()
	result.context = context.Background()
	result.configuration = configuration
	result.client = client.New(configuration)
	result.metric = metrics.New(configuration)
	result.cluster = cluster

	return result
}
