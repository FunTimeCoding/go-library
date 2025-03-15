package constant

const (
	None      = "none"
	NoComment = "no comment"

	HostEnvironment = "ALERTMANAGER_HOST"
)

const (
	NameField       = "alertname"
	SeverityField   = "severity"
	SummaryField    = "summary"
	MessageField    = "message"
	PrometheusField = "prometheus"
)

// State
const (
	ActiveState     = "active"
	ExpiredState    = "expired"
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

const NodeNotReady = "KubeNodeNotReady"

var (
	States = []string{
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
