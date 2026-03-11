package lint

import "fmt"

func stubTestContent(packageName string) string {
	return fmt.Sprintf(
		"package %s\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc TestStub(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		packageName,
	)
}
