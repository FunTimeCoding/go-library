package list

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	batch "k8s.io/api/batch/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Job(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	node string,
) []batch.Job {
	var o meta.ListOptions

	if node != constant.NodeAll {
		o.FieldSelector = fmt.Sprintf("spec.nodeName=%s", node)
	}

	result, e := operation.Job(c, namespace).List(x, o)
	errors.PanicOnError(e)

	return result.Items
}
