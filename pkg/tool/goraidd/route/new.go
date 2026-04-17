package route

func New(
	logCachePath string,
	outputPath string,
) *Router {
	return &Router{
		logCachePath: logCachePath,
		outputPath:   outputPath,
	}
}
