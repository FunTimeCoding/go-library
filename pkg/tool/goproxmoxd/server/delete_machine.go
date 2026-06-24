package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func (s *Server) DeleteMachine(
	_ context.Context,
	r server.DeleteMachineRequestObject,
) (server.DeleteMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DeleteMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.DeleteMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.DeleteMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	if vm.Status == "running" {
		return server.DeleteMachine400JSONResponse(
			*clientError(
				fmt.Errorf(
					"vm %d is running — stop it before deleting",
					r.Identifier,
				),
			),
		), nil
	}

	purge := false

	if r.Params.Purge != nil {
		purge = *r.Params.Purge
	}

	task, e := c.DeleteMachine(
		vm,
		&proxmox.VirtualMachineDeleteOptions{
			Purge:                    proxmox.IntOrBool(purge),
			DestroyUnreferencedDisks: proxmox.IntOrBool(true),
		},
	)

	if e != nil {
		return server.DeleteMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	e = c.WaitForTask(task, 120)

	if e != nil {
		return server.DeleteMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteMachine200JSONResponse{
		TaskId: "deleted",
	}, nil
}
