package client_configuration

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func path() string {
	return system.Join(system.Home(), constant.KubernetesConfigurationPath)
}
