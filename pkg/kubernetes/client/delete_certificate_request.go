package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) DeleteCertificateRequest(
	namespace string,
	name string,
) {
	errors.PanicOnError(
		c.dynamic.Resource(constant.CertificateRequestGVR).Namespace(namespace).Delete(
			c.context,
			name,
			meta.DeleteOptions{},
		),
	)
}
