package web

import "github.com/funtimecoding/go-library/pkg/raid_parser"

type Server struct {
	logCachePath      string
	eliteInsightsPath string
	outputPath        string
	parser            *raid_parser.Client
}
