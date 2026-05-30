package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetTasks(
	_ context.Context,
	r server.GetTasksRequestObject,
) (server.GetTasksResponseObject, error) {
	var taskType string

	if r.Params.Type != nil {
		taskType = string(*r.Params.Type)
	}

	tasks := convert.Tasks(s.habitica.MustTasks(taskType))
	result := make(server.GetTasks200JSONResponse, len(tasks))

	for i, t := range tasks {
		result[i] = *t
	}

	return result, nil
}
