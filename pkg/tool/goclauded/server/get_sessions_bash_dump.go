package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionsBashDump(
	w http.ResponseWriter,
	r *http.Request,
) {
	var commands []string

	for _, session := range s.claude.Sessions() {
		for _, call := range s.claude.ToolCalls(session.Identifier) {
			if call.Name == "Bash" && call.Detail != "" {
				commands = append(commands, call.Detail)
			}
		}
	}

	web.EncodeNotation(
		w,
		server.BashDumpResponse{Commands: commands},
	)
}
