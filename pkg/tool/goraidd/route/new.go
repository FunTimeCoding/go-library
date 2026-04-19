package route

import (
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
)

func New(
	s *store.Store,
	eliteInsightsPath string,
	outputPath string,
	p *raid_parser.Client,
) *Router {
	return &Router{
		store:             s,
		eliteInsightsPath: eliteInsightsPath,
		outputPath:        outputPath,
		parser:            p,
	}
}
