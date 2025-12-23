package constant

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/prometheus/common/model"
	"time"
)

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

	HighMemoryUsage = "HighMemoryUsage" // Test alert name

	Alerts = "/alerts"

	PermanentTag = "#permanent"
)

// Alert label
const (
	AlertnameLabel = model.AlertNameLabel

	SummaryLabel     = "summary"
	DescriptionLabel = "description"
	PrometheusLabel  = "prometheus"
	SeverityLabel    = "severity"
	MessageLabel     = "message"
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
	Format = option.Color.Copy().Tag(tag.Link, tag.Comment)

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
