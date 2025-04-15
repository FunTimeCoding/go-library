package issue

import (
	"github.com/andygrunwald/go-jira"
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

	CommentNameFilter []string

	ShortStatus face.StringAlias
	score       float64
	scoreColor  face.SprintFunction
	ageColor    face.SprintFunction

	Raw *jira.Issue
}
