package checksum

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"strings"
)

func Map(workDirectory string) map[string]string {
	return Parse(strings.TrimSpace(system.ReadFile(Path(workDirectory))))
}
