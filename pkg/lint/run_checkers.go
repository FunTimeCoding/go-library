package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"strings"
)

func runCheckers(
	v *virtual_file_system.System,
	fixes *virtual_file_system.System,
	paths []string,
	checkers []Checker,
	fix bool,
	verbose bool,
	r *output.Results,
) {
	for _, p := range paths {
		if verbose {
			fmt.Printf("Process: %s\n", p)
		}

		original := v.Read(p)
		content := original

		for _, check := range checkers {
			result := check(p, strings.NewReader(content))

			for _, c := range result.Concerns {
				if c.Fixed && fix {
					r.Add(p, concernMessage(c))
				} else if !c.Fixed {
					r.AddBlocked(p, c.Text)
				}
			}

			if result.Fixed != "" && result.Fixed != result.Original {
				content = result.Fixed
			}
		}

		if content != original {
			fixes.Write(p, content)
		}
	}
}
