package scan

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/parse"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
)

func hasSentryReporter(
	v *virtual_file_system.System,
	directory string,
) bool {
	for _, name := range v.MustReadDirectory(directory) {
		if !strings.HasSuffix(name, constant.GoExtension) {
			continue
		}

		f, _, e := parse.Source(
			name,
			v.ReadString(filepath.Join(directory, name)),
		)

		if e != nil {
			continue
		}

		if parse.HasCall(
			f,
			"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter",
			"New",
		) {
			return true
		}
	}

	return false
}
