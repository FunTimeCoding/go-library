package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/detail"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/user_map"
	"github.com/funtimecoding/go-library/pkg/face"
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
	Description     string

	Team              *team.Team
	TeamKey           string
	TeamMap           *team_map.Map
	UserMap           *user_map.Map
	Prometheus        *detail.Prometheus
	shortUser         face.StringAlias
	shortAlert        face.StringAlias
	descriptionToName face.StringAlias
	parseDescription  constant.ParseDescription
	MonitorIdentifier string

	Raw       *alert.Alert
	RawResult *alert.GetAlertResult
}
