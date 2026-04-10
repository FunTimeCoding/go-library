package firefox

import "github.com/funtimecoding/go-library/pkg/firefox/constant"

func New(o ...Option) *Client {
	result := &Client{
		address: constant.DefaultHost,
		pending: make(map[int]chan response),
	}

	for _, f := range o {
		f(result)
	}

	return result
}
