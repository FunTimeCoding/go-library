package client_configuration

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewContext(context string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{
			ExplicitPath: system.Join(
				system.Home(),
				constant.KubernetesConfigurationPath,
			),
		},
		&clientcmd.ConfigOverrides{CurrentContext: context},
	).ClientConfig()
}
