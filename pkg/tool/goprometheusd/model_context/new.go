package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/service"
)

func New(
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		service:  v,
		reporter: r,
	}
	result.register()

	return result
}
