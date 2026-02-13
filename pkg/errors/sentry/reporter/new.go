package reporter

func New(
	project string,
	locator string,
	environment string,
	version string,
) *Reporter {
	return &Reporter{
		project:     project,
		locator:     locator,
		environment: environment,
		version:     version,
	}
}
