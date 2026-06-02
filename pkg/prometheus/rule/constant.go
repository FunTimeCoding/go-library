package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

const (
	SeverityKey = "severity"

	SummaryKey       = "summary"
	DescriptionKey   = "description"
	DurationKey      = "duration"
	RunbookKey       = "runbook_url"
	DocumentationKey = "docs"

	AlertType   = "alert"
	RecordType  = "record"
	AllType     = "all"
	UnknownType = "unknown"
)

const (
	HealthOkay v1.RuleHealth = "ok"

	InactiveState = "inactive"
	PendingState  = "pending"
	FiringState   = "firing"
)

var (
	Healths = []v1.RuleHealth{
		HealthOkay,
	}
	States = []string{
		InactiveState,
		PendingState,
		FiringState,
	}
)
