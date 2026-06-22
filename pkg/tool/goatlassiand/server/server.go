package server

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/face"
)

type Server struct {
	jira       *jira.Client
	confluence *confluence.Client
	reporter   face.Reporter
}
