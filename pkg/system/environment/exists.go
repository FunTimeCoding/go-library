package environment

import "os"

func Exists(name string) bool {
	_, exists := os.LookupEnv(name)

	return exists
}
