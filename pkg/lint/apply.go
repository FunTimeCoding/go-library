package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"github.com/funtimecoding/go-library/pkg/system/vfs"
)

func apply(
	fixes *vfs.VFS,
	path string,
	r *file_report.Report,
) {
	for _, c := range r.Concerns {
		fmt.Printf("%s: %s\n", c.Text, c.Path)
	}

	if r.Fixed != "" && r.Fixed != r.Original {
		fixes.Write(path, r.Fixed)
	}
}
