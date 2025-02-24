package build

import "fmt"

func SystemArchitecture(
	operatingSystem string,
	architecture string,
) string {
	return fmt.Sprintf("%s-%s", operatingSystem, architecture)
}
