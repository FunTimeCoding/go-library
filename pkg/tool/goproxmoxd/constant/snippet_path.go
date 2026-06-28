package constant

import "fmt"

func SnippetPath(name string) string {
	return fmt.Sprintf("%s/%s", SnippetDirectory, name)
}
