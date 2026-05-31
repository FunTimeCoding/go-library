package goclaude

import "fmt"

func printPeekEntry(
	userText string,
	assistantContext *string,
	depth int,
) {
	fmt.Println(userText)

	if depth > 1 && assistantContext != nil && *assistantContext != "" {
		text := *assistantContext
		limit := (depth - 1) * 100

		if len(text) > limit {
			text = text[:limit]
		}

		fmt.Printf("  → %s\n", text)
	}
}
