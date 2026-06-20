package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/provision/store"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/runner"
)

func New(
	n *runner.Runner,
	s *store.Store,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		runner:   n,
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
