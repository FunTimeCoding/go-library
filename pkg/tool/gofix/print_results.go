package gofix

import "fmt"

func printResults(entries []result, summary bool) bool {
	var hasBlocked bool
	seen := make(map[string]bool)

	for _, e := range entries {
		if e.blocked {
			hasBlocked = true
			fmt.Printf("%s: %s\n", e.path, e.message)

			continue
		}

		if summary {
			if !seen[e.path] {
				seen[e.path] = true
				fmt.Println(e.path)
			}

			continue
		}

		fmt.Printf("%s: %s\n", e.path, e.message)
	}

	return hasBlocked
}
