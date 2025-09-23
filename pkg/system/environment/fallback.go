package environment

import "os"

func Fallback(
	name string,
	fallback string,
) string {
	result := os.Getenv(name)

	if result == "" {
		return fallback
	}

	return result
}
