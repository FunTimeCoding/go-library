package route

func New(
	parserPath string,
	templatePath string,
	outputPath string,
) *Router {
	return &Router{
		parserPath:   parserPath,
		templatePath: templatePath,
		outputPath:   outputPath,
	}
}
