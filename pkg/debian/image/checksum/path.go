package checksum

import "github.com/funtimecoding/go-library/pkg/system"

func Path(workDirectory string) string {
	return system.Join(workDirectory, File)
}
