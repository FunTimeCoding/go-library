package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) RemoveVirtualTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	vm, e := s.client.RemoveVirtualTag(name, tag)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	entry := server.VirtualMachine{Identifier: vm.Identifier, Name: vm.Name}

	if len(vm.Tags) > 0 {
		entry.Tags = &vm.Tags
	}

	web.EncodeNotation(w, entry)
}
