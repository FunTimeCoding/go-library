package alert

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/alert/prometheus_detail"
	"github.com/funtimecoding/go-library/pkg/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/opsgenie/user_map"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"time"
)

type Alert struct {
	Identifier      string
	SmallIdentifier int
	Name            string
	Status          string
	Seen            bool
	Acknowledged    bool
	Snoozed         bool
	Priority        alert.Priority
	Link            string
	Owner           string
	Source          string
	CreatedAt       time.Time
	SnoozedUntil    time.Time
	UpdatedAt       time.Time
	Report          alert.Report
	Responders      []alert.Responder
	Tags            []string
	Details         map[string]string

	Team              *team.Team
	TeamKey           string
	TeamMap           *team_map.Map
	UserMap           *user_map.Map
	Prometheus        *prometheus_detail.Detail
	shortUser         constant.StringAlias
	shortAlert        constant.StringAlias
	descriptionToName constant.StringAlias
	MonitorIdentifier string

	Raw       *alert.Alert
	RawResult *alert.GetAlertResult
}
