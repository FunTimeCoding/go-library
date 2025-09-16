package environment

import "os"

func Default(
	name string,
	fallback string,
) string {
	result := os.Getenv(name)

	if result == "" {
		return fallback
	}

	return result
}
