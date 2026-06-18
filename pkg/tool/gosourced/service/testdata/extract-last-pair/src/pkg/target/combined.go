package target

import "fmt"

func FormatName(name string) string {
	return fmt.Sprintf("name: %s", name)
}

func SummarizeName(name string) string {
	return fmt.Sprintf("[%s]", name)
}
