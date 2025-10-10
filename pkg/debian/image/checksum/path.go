package checksum

import "github.com/funtimecoding/go-library/pkg/system/join"

func Path(workDirectory string) string {
	return join.Absolute(workDirectory, File)
}
