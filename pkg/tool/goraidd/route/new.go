package route

import "github.com/funtimecoding/go-library/pkg/raid_parser"

func New(
	logCachePath string,
	eliteInsightsPath string,
	outputPath string,
	parser *raid_parser.Client,
) *Router {
	return &Router{
		logCachePath:      logCachePath,
		eliteInsightsPath: eliteInsightsPath,
		outputPath:        outputPath,
		parser:            parser,
	}
}
