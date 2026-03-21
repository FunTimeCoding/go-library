package lint

import "strings"

func stubTestName(packageName string) string {
	if packageName == "main" {
		return "Stub"
	}

	suffix := stubTestSuffix(packageName)

	return strings.ToUpper(suffix[:1]) + suffix[1:]
}
