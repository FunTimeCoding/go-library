package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func (s *Server) DeleteMachine(
	_ context.Context,
	r server.DeleteMachineRequestObject,
) (server.DeleteMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteMachine400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DeleteMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.DeleteMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if vm == nil {
		return server.DeleteMachine404Response{}, nil
	}

	if vm.Status == "running" {
		return server.DeleteMachine400JSONResponse{
			ClientErrorJSONResponse: *clientError(
				fmt.Errorf(
					"vm %d is running — stop it before deleting",
					r.Identifier,
				),
			),
		}, nil
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
		return server.DeleteMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	e = c.WaitForTask(task, 120)

	if e != nil {
		return server.DeleteMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	return server.DeleteMachine200JSONResponse{
		TaskId: "deleted",
	}, nil
}
