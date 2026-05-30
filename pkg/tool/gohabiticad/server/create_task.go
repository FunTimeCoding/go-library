package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) CreateTask(
	_ context.Context,
	r server.CreateTaskRequestObject,
) (server.CreateTaskResponseObject, error) {
	c := request.New()
	c.Type = string(r.Body.Type)
	c.Text = r.Body.Text

	if r.Body.Notes != nil {
		c.Notes = *r.Body.Notes
	}

	return server.CreateTask201JSONResponse(
		*convert.Task(s.habitica.MustCreateTask(c)),
	), nil
}
