package web

func NewServer(
	logCachePath string,
	eliteInsightsPath string,
	outputPath string,
) *Server {
	return &Server{
		logCachePath:      logCachePath,
		eliteInsightsPath: eliteInsightsPath,
		outputPath:        outputPath,
	}
}
