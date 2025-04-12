package environment

import "os"

func GetDefault(
	name string,
	fallback string,
) string {
	result := os.Getenv(name)

	if result == "" {
		return fallback
	}

	return result
}
