package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	var contexts []string

	if s := environment.GetDefault(
		constant.ContextEnvironment,
		"",
	); s != "" {
		contexts = split.Comma(s)
	}

	return New(contexts)
}
