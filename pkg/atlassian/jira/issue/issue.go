package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/go-library/pkg/face"
	"time"
)

type Issue struct {
	MonitorIdentifier string
	Key               string
	Summary           string
	Description       string
	Initials          string
	Status            string
	Type              string
	Create            *time.Time
	Link              string
	Labels            []string

	concerns          []string
	score             float64
	commentNameFilter []string
	shortStatus       face.StringAlias
	scoreColor        face.SprintFunction
	ageColor          face.SprintFunction
	fieldMap          *field_map.Map
	option            *option.Issue

	Raw *jira.Issue
}
