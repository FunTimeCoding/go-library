package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/constant"
)

func New(
	c *chromium.Client,
	downloadDirectory string,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:            c,
		downloadDirectory: downloadDirectory,
		snapshotCache:     make(map[string]map[string]int64),
		reporter:          r,
	}
	result.register()

	return result
}
