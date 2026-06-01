package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
)

func New(
	s *store.Store,
	n face.EventNotifier,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		store:    s,
		notifier: n,
		reporter: r,
	}
	result.register()

	return result
}
