package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetTasks(
	w http.ResponseWriter,
	_ *http.Request,
	p server.GetTasksParams,
) {
	var taskType string

	if p.Type != nil {
		taskType = string(*p.Type)
	}

	web.EncodeNotation(w, convert.Tasks(s.habitica.Tasks(taskType)))
}
