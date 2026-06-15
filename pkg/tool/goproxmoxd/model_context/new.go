package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/service"
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
		).WithInstructions(
			constant.Identity.RenderInstructions(
				map[string]bool{
					constant.MultiInstance: len(v.Instances()) > 1,
				},
			),
		).WithRecorder(t).Server(),
		service:  v,
		reporter: r,
	}
	result.register()

	return result
}
