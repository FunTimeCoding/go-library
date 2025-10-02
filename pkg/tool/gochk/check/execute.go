package check

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"strings"
)

func Execute(c []string) string {
	return strings.TrimSpace(run.New().Start(c...))
}
