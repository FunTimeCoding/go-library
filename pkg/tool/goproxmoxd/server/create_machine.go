package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateMachine(
	_ context.Context,
	r server.CreateMachineRequestObject,
) (server.CreateMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.CreateMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	m := convertCreateMachine(r.Body)
	identifier := m.Identifier

	if identifier <= 0 {
		identifier, e = c.NextIdentifier()

		if e != nil {
			return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
		}
	}

	node, e := c.Node(r.Body.Node)

	if e != nil {
		return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	options := m.BuildOptions()
	task, e := c.CreateMachine(node, identifier, options...)

	if e != nil {
		return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	e = c.WaitForTask(task, 300)

	if e != nil {
		return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	if m.Start {
		vm, e := c.Machine(node, identifier)

		if e != nil {
			return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
		}

		startTask, e := c.StartMachine(vm)

		if e != nil {
			return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
		}

		e = c.WaitForTask(startTask, 120)

		if e != nil {
			return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
		}
	}

	status := "created"

	return server.CreateMachine200JSONResponse{
		Identifier: int64(identifier),
		Status:     &status,
	}, nil
}
