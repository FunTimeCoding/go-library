package argument

import "fmt"

func CobraVersion(
	version string,
	gitHash string,
	buildDate string,
) string {
	return fmt.Sprintf("%s (%s %s)", version, gitHash, buildDate)
}
