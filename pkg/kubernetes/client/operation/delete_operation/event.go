package delete_operation

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Event(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	name string,
) {
	errors.PanicOnError(
		operation.Event(c, namespace).Delete(
			x,
			name,
			meta.DeleteOptions{},
		),
	)
}
