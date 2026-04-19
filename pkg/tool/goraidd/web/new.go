package web

import (
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
)

func NewServer(
	store *store.Store,
	eliteInsightsPath string,
	outputPath string,
	parser *raid_parser.Client,
) *Server {
	return &Server{
		store:             store,
		eliteInsightsPath: eliteInsightsPath,
		outputPath:        outputPath,
		parser:            parser,
	}
}
