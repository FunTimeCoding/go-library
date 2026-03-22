package formula

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"strings"
)

func (f *Formula) Format(o *option.Format) string {
	s := status.New(o).String(
		f.Name,
		fmt.Sprintf(
			"%s → %s",
			strings.Join(f.InstalledVersions, ", "),
			f.CurrentVersion,
		),
	)
	s.DetailLink(f.Link, "Homebrew", "")

	return s.Format()
}
