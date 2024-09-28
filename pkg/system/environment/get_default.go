package environment

import "os"

func GetDefault(
	name string,
	defaultValue string,
) string {
	result := os.Getenv(name)

	if result == "" {
		return defaultValue
	}

	return result
}
