package lint

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"path/filepath"
	"sort"
	"strings"
)

func missingSentryPrograms(
	v *virtual_file_system.System,
) []string {
	programs := make(map[string]bool)

	for _, p := range v.Files() {
		d := filepath.Dir(p)

		if !strings.HasPrefix(d, "cmd/") {
			continue
		}

		parts := strings.SplitN(d, separator.Slash, 3)

		if len(parts) < 2 {
			continue
		}

		name := parts[1]

		if name == "example" {
			continue
		}

		if _, ok := programs[name]; !ok {
			programs[name] = false
		}
	}

	for _, p := range v.Files() {
		if !strings.HasPrefix(p, "pkg/tool/") {
			continue
		}

		if !strings.HasSuffix(p, ".go") {
			continue
		}

		content := v.Read(p)

		if strings.Contains(content, "reporter.New(") {
			d := filepath.Dir(p)
			parts := strings.SplitN(d, separator.Slash, 4)

			if len(parts) >= 3 {
				name := parts[2]

				if _, ok := programs[name]; ok {
					programs[name] = true
				}
			}
		}
	}

	var result []string

	for name, hasSentry := range programs {
		if !hasSentry {
			result = append(result, name)
		}
	}

	sort.Strings(result)

	return result
}
