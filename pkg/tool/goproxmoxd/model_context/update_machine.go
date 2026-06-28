package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/luthermonson/go-proxmox"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) UpdateMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateMachine,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	var setFields []proxmox.VirtualMachineOption
	var deleteFields []string
	setNames := make(map[string]bool)

	if a.Name != "" {
		setFields = append(
			setFields,
			proxmox.VirtualMachineOption{Name: "name", Value: a.Name},
		)
		setNames["name"] = true
	}

	if a.Tags != "" {
		setFields = append(
			setFields,
			proxmox.VirtualMachineOption{Name: "tags", Value: a.Tags},
		)
		setNames["tags"] = true
	}

	if a.OnBoot != nil {
		value := 0

		if *a.OnBoot {
			value = 1
		}

		setFields = append(
			setFields,
			proxmox.VirtualMachineOption{Name: "onboot", Value: value},
		)
		setNames["onboot"] = true
	}

	if a.Cores > 0 {
		setFields = append(
			setFields,
			proxmox.VirtualMachineOption{Name: "cores", Value: a.Cores},
		)
		setNames["cores"] = true
	}

	if a.Memory > 0 {
		setFields = append(
			setFields,
			proxmox.VirtualMachineOption{Name: "memory", Value: a.Memory},
		)
		setNames["memory"] = true
	}

	if a.Description != "" {
		setFields = append(
			setFields,
			proxmox.VirtualMachineOption{
				Name:  "description",
				Value: a.Description,
			},
		)
		setNames["description"] = true
	}

	if a.Delete != "" {
		for _, field := range strings.Split(a.Delete, ",") {
			name := strings.TrimSpace(field)

			if setNames[name] {
				return response.Fail(
					"cannot set and delete %q in the same call",
					name,
				)
			}

			deleteFields = append(deleteFields, name)
		}
	}

	if len(setFields) == 0 && len(deleteFields) == 0 {
		return response.Fail("no changes specified")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	vm, e := findMachine(c, a.Identifier, a.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return response.Fail("%s", e)
		}

		return s.captureDetail(e)
	}

	if len(deleteFields) > 0 {
		setFields = append(
			setFields,
			proxmox.VirtualMachineOption{
				Name:  "delete",
				Value: strings.Join(deleteFields, ","),
			},
		)
	}

	e = c.UpdateMachineConfiguration(vm, setFields...)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("updated")
}
