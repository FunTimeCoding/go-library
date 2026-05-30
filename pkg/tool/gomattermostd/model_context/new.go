package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/monitor"
)

func New(
	m *mattermost.Client,
	o *monitor.Monitor,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:   m,
		monitor:  o,
		reporter: r,
	}
	result.register()

	return result
}
