package route

import (
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
)

type Router struct {
	store             *store.Store
	eliteInsightsPath string
	outputPath        string
	parser            *raid_parser.Client
}
