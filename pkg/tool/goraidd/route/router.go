package route

import "github.com/funtimecoding/go-library/pkg/raid_parser"

type Router struct {
	logCachePath      string
	eliteInsightsPath string
	outputPath        string
	parser            *raid_parser.Client
}
