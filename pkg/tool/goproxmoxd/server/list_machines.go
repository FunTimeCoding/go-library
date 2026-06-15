package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListMachines(
	_ context.Context,
	r server.ListMachinesRequestObject,
) (server.ListMachinesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListMachines400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListMachines500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	var pointers []*server.Machine

	if r.Params.Node != nil && *r.Params.Node != "" {
		node, e := c.Node(*r.Params.Node)

		if e != nil {
			return server.ListMachines500JSONResponse{
				ErrorJSONResponse: *s.captureFail(
					e,
					constant.UnexpectedError,
				),
			}, nil
		}

		machines, e := c.Machines(node)

		if e != nil {
			return server.ListMachines500JSONResponse{
				ErrorJSONResponse: *s.captureFail(
					e,
					constant.UnexpectedError,
				),
			}, nil
		}

		pointers = convert.Machines(machines)
	} else {
		nodes, e := c.Nodes()

		if e != nil {
			return server.ListMachines500JSONResponse{
				ErrorJSONResponse: *s.captureFail(
					e,
					constant.UnexpectedError,
				),
			}, nil
		}

		for _, ns := range nodes {
			node, e := c.Node(ns.Node)

			if e != nil {
				continue
			}

			machines, e := c.Machines(node)

			if e != nil {
				continue
			}

			pointers = append(pointers, convert.Machines(machines)...)
		}
	}

	result := make(server.ListMachines200JSONResponse, len(pointers))

	for i, m := range pointers {
		result[i] = *m
	}

	return result, nil
}
