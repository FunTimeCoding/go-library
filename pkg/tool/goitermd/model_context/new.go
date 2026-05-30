package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/goitermd/constant"
	iterm "github.com/funtimecoding/go-library/pkg/tool/goitermd/face"
)

func New(
	c iterm.ItermSource,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:   c,
		reporter: r,
	}
	result.register()

	return result
}
