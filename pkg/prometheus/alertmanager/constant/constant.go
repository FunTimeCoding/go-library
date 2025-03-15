package constant

const (
	None      = "none"
	NoComment = "no comment"

	HostEnvironment = "ALERTMANAGER_HOST"
)

// Alert field
const (
	NameField       = "alertname"
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
