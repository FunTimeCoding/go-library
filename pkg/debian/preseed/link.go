package preseed

import "fmt"

func Link(release string) string {
	return fmt.Sprintf(
		"https://www.debian.org/releases/%s/example-preseed.txt",
		release,
	)
}
