package goanalyze

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"regexp"
)

func fixFileReferences(
	path string,
	renames []exportedRename,
) {
	content, e := os.ReadFile(path)

	if e != nil {
		return
	}

	modified := content

	for _, r := range renames {
		pattern := regexp.MustCompile(`\b` + regexp.QuoteMeta(r.oldName) + `\b`)
		replaced := pattern.ReplaceAll(modified, []byte(r.newName))

		if !bytes.Equal(replaced, modified) {
			fmt.Printf(
				"Renamed: %s → %s (unloaded: %s)\n",
				r.oldName,
				r.newName,
				path,
			)
			modified = replaced
		}
	}

	if !bytes.Equal(modified, content) {
		errors.PanicOnError(os.WriteFile(path, modified, 0644))
	}
}
