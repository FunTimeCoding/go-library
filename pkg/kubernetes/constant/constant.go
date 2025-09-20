package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	ContextEnvironment     = "KUBERNETES_CONTEXT"
	AutoCleanupEnvironment = "KUBERNETES_AUTO_CLEANUP"
)

// NodeAll analogous to NamespaceAll
const NodeAll = ""

const (
	Kubectl        = "kubectl"
	Configuration  = "config"
	CurrentContext = "current-context"
)

var (
	Format = option.ExtendedColor
	Dense  = option.Color
)

const (
	PodsResource       = "pods"
	ExecuteSubResource = "exec"
)

const DNSConfigForming = "DNSConfigForming" // Event reason

var IrrelevantEventReason = []string{
	DNSConfigForming,
}

const (
	TrivyNamespace = "trivy"
	TrivyCron      = "trivy"

	RenovateNamespace = "renovate"
	RenovateCron      = "renovate"

	ManualCron = "manual"
)
