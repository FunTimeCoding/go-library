package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const ContextEnvironment = "KUBERNETES_CONTEXT"

// NodeAll analogous to NamespaceAll
const NodeAll = ""

const (
	Kubectl        = "kubectl"
	Configuration  = "config"
	CurrentContext = "current-context"
)

var Format = option.ExtendedColor
