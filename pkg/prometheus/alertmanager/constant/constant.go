package constant

import "time"

const (
	None        = "none"
	NoComment   = "no comment"
	UnknownRule = "unknown rule"

	DefaultDuration = 10 * time.Minute

	HostEnvironment     = "ALERTMANAGER_HOST"
	UserEnvironment     = "ALERTMANAGER_USER"
	PasswordEnvironment = "ALERTMANAGER_PASSWORD"

	AmtoolCommand = "amtool"

	AmtoolPath          = "amtool"
	AmtoolConfiguration = "config.yml"

	AmtoolConfigurationPrefix = "config-"

	KubernetesPrefix = "Kube"
)

// Alert field
const (
	AlertnameField  = "alertname"
	SeverityField   = "severity"
	SummaryField    = "summary"
	MessageField    = "message"
	PrometheusField = "prometheus"
)

// Alert state
const (
	ActiveState     = "active"
	SuppressedState = "suppressed"
)

// Severity
const (
	CriticalSeverity = "critical"
	InfoSeverity     = "info"
	NoneSeverity     = "none"
	UnknownSeverity  = "unknown"
	WarningSeverity  = "warning"
)

const ExpiredState = "expired" // Silence state

const NodeNotReady = "KubeNodeNotReady" // Alert name

var (
	AlertStates = []string{
		ActiveState,
		SuppressedState,
	}
	Severities = []string{
		CriticalSeverity,
		WarningSeverity,
		InfoSeverity,
		NoneSeverity,
		UnknownSeverity,
	}
	SevereSeverities = []string{
		CriticalSeverity,
		WarningSeverity,
	}
	RedSeverities = []string{
		CriticalSeverity,
	}
	YellowSeverities = []string{
		WarningSeverity,
	}
)
