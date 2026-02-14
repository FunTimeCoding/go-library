package main

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/maintenance_log"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	s := server.New(maintenance_log.New().Nested())
	s.Start()
	system.KillSignalBlock()
	s.Stop()
}
