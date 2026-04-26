package string_constant

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"os"
	"path/filepath"
	"strings"
)

func collectFromConstantDirectory(
	result map[string]knownConstant,
	directory string,
	subDirectory string,
	distance int,
) {
	constantDirectory := filepath.Join(directory, subDirectory)
	entries, e := os.ReadDir(constantDirectory)

	if e != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(
			entry.Name(),
			constant.GoExtension,
		) {
			continue
		}

		if strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}

		parseConstants(
			result,
			filepath.Join(constantDirectory, entry.Name()),
			subDirectory,
			distance,
		)
	}
}
