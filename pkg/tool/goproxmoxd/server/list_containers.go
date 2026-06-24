package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListContainers(
	_ context.Context,
	r server.ListContainersRequestObject,
) (server.ListContainersResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListContainers400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListContainers500JSONResponse(*s.captureDetail(e)), nil
	}

	var pointers []*server.Container

	if r.Params.Node != nil && *r.Params.Node != "" {
		node, e := c.Node(*r.Params.Node)

		if e != nil {
			return server.ListContainers500JSONResponse(*s.captureDetail(e)), nil
		}

		containers, e := c.Containers(node)

		if e != nil {
			return server.ListContainers500JSONResponse(*s.captureDetail(e)), nil
		}

		pointers = convert.Containers(containers)
	} else {
		nodes, e := c.Nodes()

		if e != nil {
			return server.ListContainers500JSONResponse(*s.captureDetail(e)), nil
		}

		for _, ns := range nodes {
			node, e := c.Node(ns.Node)

			if e != nil {
				continue
			}

			containers, e := c.Containers(node)

			if e != nil {
				continue
			}

			pointers = append(pointers, convert.Containers(containers)...)
		}
	}

	result := make(server.ListContainers200JSONResponse, len(pointers))

	for i, ct := range pointers {
		result[i] = *ct
	}

	return result, nil
}
