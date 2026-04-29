package server

import (
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
)

func New(
	s *store.Store,
	elitePath string,
	outputPath string,
	p *raid_parser.Client,
) *Server {
	return &Server{
		store:      s,
		elitePath:  elitePath,
		outputPath: outputPath,
		parser:     p,
	}
}
