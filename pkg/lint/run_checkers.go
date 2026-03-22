package lint

import (
	"fmt"
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
) {
	for _, p := range paths {
		if verbose {
			fmt.Printf("Process: %s\n", p)
		}

		original := v.Read(p)
		content := original

		for _, check := range checkers {
			r := check(p, strings.NewReader(content))

			for _, c := range r.Concerns {
				if c.Fixed && fix {
					fmt.Printf("%s: %s (auto-fixed)\n", c.Text, c.Path)
				} else if c.Fixed {
					fmt.Printf("%s: %s (auto-fixable)\n", c.Text, c.Path)
				} else {
					fmt.Printf("%s: %s\n", c.Text, c.Path)
				}
			}

			if r.Fixed != "" && r.Fixed != r.Original {
				content = r.Fixed
			}
		}

		if content != original {
			fixes.Write(p, content)
		}
	}
}
