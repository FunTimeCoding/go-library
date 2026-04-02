package reporter

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"os"
)

func (r *Reporter) RecoverFlush(v any) {
	r.hub.Recover(v)
	sentry.Flush(r.hub)

	if v != nil {
		fmt.Printf("Captured panic: %v\n", v)
		os.Exit(1)
	}
}
