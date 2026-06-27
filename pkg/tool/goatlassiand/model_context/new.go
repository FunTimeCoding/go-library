package model_context

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	atlassianFace "github.com/funtimecoding/go-library/pkg/tool/goatlassiand/face"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/service"
)

func New(
	j *jira.Client,
	c atlassianFace.ConfluenceSource,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		jira:       j,
		confluence: c,
		service:    service.New(c),
		reporter:   r,
	}
	result.register()

	return result
}
