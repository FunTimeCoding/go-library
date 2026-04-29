package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) RemoveVirtualTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	vm := s.client.RemoveVirtualTag(name, tag)
	entry := server.VirtualMachine{Identifier: vm.Identifier, Name: vm.Name}

	if len(vm.Tags) > 0 {
		entry.Tags = &vm.Tags
	}

	web.EncodeNotation(w, entry)
}
