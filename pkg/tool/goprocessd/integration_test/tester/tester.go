package tester

import "github.com/funtimecoding/go-library/pkg/tool/goprocessd/server"

type Tester struct {
	Server       *server.Server
	Directory    string
	SocketPath   string
	ProcfilePath string
	EnvrcPath    string
}
