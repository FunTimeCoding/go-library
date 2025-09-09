package check

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"strings"
)

func RunCommand(c []string) string {
	return strings.TrimSpace(run.New().Start(c...))
}
