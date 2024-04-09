package gitlab

import (
	"github.com/xanzy/go-gitlab"
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
