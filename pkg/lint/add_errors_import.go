package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"strings"
)

func addErrorsImport(content string) string {
	if strings.Contains(content, "import (") {
		return strings.Replace(
			content,
			"import (",
			"import (\n\t"+constant.ErrorsImportPath,
			1,
		)
	}

	singleImport := "import "
	index := strings.Index(content, singleImport)

	if index == -1 {
		return content
	}

	lineEnd := strings.Index(content[index:], "\n")

	if lineEnd == -1 {
		return content
	}

	existingImport := content[index : index+lineEnd]

	return strings.Replace(
		content,
		existingImport,
		"import (\n\t"+constant.ErrorsImportPath+"\n\t"+
			strings.TrimPrefix(existingImport, "import ")+"\n)",
		1,
	)
}
