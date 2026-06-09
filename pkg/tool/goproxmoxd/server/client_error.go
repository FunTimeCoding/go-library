package server

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"

func clientError(e error) *server.ClientErrorJSONResponse {
	return &server.ClientErrorJSONResponse{Error: e.Error()}
}
