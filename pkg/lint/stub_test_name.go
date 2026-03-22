package lint

import (
	"fmt"
	"strings"
)

func stubTestName(packageName string) string {
	if packageName == "main" {
		return "Stub"
	}

	suffix := stubTestSuffix(packageName)

	return fmt.Sprintf("%s%s", strings.ToUpper(suffix[:1]), suffix[1:])
}
