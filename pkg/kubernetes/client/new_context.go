package client

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/client_configuration"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/metrics"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewContext(cluster string) (*Client, error) {
	c, configurationFail := client_configuration.NewContext(cluster)

	if configurationFail != nil {
		return nil, configurationFail
	}

	d, dynamicFail := client.NewDynamic(c)

	if dynamicFail != nil {
		return nil, dynamicFail
	}

	result := Stub()
	result.context = context.Background()
	result.client = client.New(c)
	result.metric = metrics.New(c)
	result.cluster = cluster
	result.dynamic = d
	_, nodesFail := operation.Node(result.client).List(
		result.context,
		meta.ListOptions{},
	)

	if nodesFail != nil {
		return nil, nodesFail
	}

	return result, nil
}
