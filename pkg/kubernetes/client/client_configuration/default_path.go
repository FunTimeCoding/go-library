package client_configuration

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func path() string {
	return join.Absolute(system.Home(), constant.KubernetesConfigurationPath)
}
