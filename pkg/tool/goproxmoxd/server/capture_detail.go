package server

import (
	"errors"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if errors.Is(e, proxmox.ErrNotFound) ||
		errors.Is(e, proxmox.ErrNotAuthorized) ||
		errors.Is(e, proxmox.ErrTimeout) {
		return s.captureFail(e, e.Error())
	}

	return s.captureFail(e, constant.UnexpectedError)
}
