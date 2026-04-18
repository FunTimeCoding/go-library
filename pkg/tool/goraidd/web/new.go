package web

import "github.com/funtimecoding/go-library/pkg/raid_parser"

func NewServer(
	logCachePath string,
	eliteInsightsPath string,
	outputPath string,
	parser *raid_parser.Client,
) *Server {
	return &Server{
		logCachePath:      logCachePath,
		eliteInsightsPath: eliteInsightsPath,
		outputPath:        outputPath,
		parser:            parser,
	}
}
