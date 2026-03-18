package lint

import (
	"fmt"
	"strings"
)

func stubTestSuffix(packageName string) string {
	parts := strings.Split(packageName, "_")

	return parts[len(parts)-1]
}

func stubTestName(packageName string) string {
	if packageName == "main" {
		return "Stub"
	}

	suffix := stubTestSuffix(packageName)

	return strings.ToUpper(suffix[:1]) + suffix[1:]
}

func stubTestContent(packageName string, testdata bool) string {
	testName := stubTestName(packageName)

	if testdata {
		return fmt.Sprintf(
			"package %s\n\nimport \"testing\"\n\nfunc Test%s(t *testing.T) {\n\tt.Helper()\n}\n",
			packageName,
			testName,
		)
	}

	return fmt.Sprintf(
		"package %s\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc Test%s(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		packageName,
		testName,
	)
}
