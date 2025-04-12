package client

import "github.com/funtimecoding/go-library/pkg/errors"

func NewContextStrict(cluster string) *Client {
	result, e := NewContext(cluster)
	errors.PanicOnError(e)

	return result
}
