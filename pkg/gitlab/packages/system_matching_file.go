package packages

import (
	"gitlab.com/gitlab-org/api/client-go"
	"runtime"
	"strings"
)

func SystemMatchingFile(v []*gitlab.PackageFile) *gitlab.PackageFile {
	for _, element := range v {
		if strings.Contains(element.FileName, runtime.GOARCH) &&
			strings.Contains(element.FileName, runtime.GOOS) {
			return element
		}
	}

	return nil
}
