package client_context

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"strings"
)

func Current() string {
	return strings.TrimSpace(
		system.Run(
			constant.Kubectl,
			constant.Configuration,
			constant.CurrentContext,
		),
	)
}
