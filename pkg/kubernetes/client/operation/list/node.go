package list

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Node(
	c *kubernetes.Clientset,
	x context.Context,
) []core.Node {
	result, e := operation.Node(c).List(x, meta.ListOptions{})
	errors.PanicOnError(e)

	return result.Items
}
