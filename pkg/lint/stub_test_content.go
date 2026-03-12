package lint

import (
	"fmt"
	"strings"
)

func stubTestSuffix(packageName string) string {
	parts := strings.Split(packageName, "_")

	return parts[len(parts)-1]
}

func stubTestContent(packageName string) string {
	suffix := stubTestSuffix(packageName)
	testName := strings.ToUpper(suffix[:1]) + suffix[1:]

	return fmt.Sprintf(
		"package %s\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc Test%s(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		packageName,
		testName,
	)
}
