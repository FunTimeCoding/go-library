package output

import "fmt"

func PrintResults(
	entries []Result,
	summary bool,
) bool {
	var hasBlocked bool
	seen := make(map[string]bool)

	for _, e := range entries {
		if e.Blocked {
			hasBlocked = true
			fmt.Printf("%s: %s\n", e.Path, e.Message)

			continue
		}

		if summary {
			if !seen[e.Path] {
				seen[e.Path] = true
				fmt.Println(e.Path)
			}

			continue
		}

		fmt.Printf("%s: %s\n", e.Path, e.Message)
	}

	return hasBlocked
}
