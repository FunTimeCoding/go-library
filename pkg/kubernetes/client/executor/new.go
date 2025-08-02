package executor

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

func New(
	c *rest.Config,
	r *rest.Request,
) remotecommand.Executor {
	result, e := remotecommand.NewSPDYExecutor(
		c,
		web.PostMethod,
		r.URL(),
	)
	errors.PanicOnError(e)

	return result
}
