package formula

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (f *Formula) Format(o *option.Format) string {
	s := status.New(o).String(
		f.Name,
		fmt.Sprintf(
			"%s → %s",
			join.CommaSpace(f.InstalledVersions),
			f.CurrentVersion,
		),
	)
	s.DetailLink(f.Link, "Homebrew", "")

	return s.Format()
}
