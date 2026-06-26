package model_context

import "fmt"

func snippetPath(name string) string {
	return fmt.Sprintf("%s/%s", snippetDirectory, name)
}
